package contr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type AiAgentAttachment struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}

type AiAgentTool struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]any         `json:"parameters"`
	Dangerous   bool                   `json:"dangerous"`
	Extra       map[string]interface{} `json:"-"`
}

type AiAgentMenuItem struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Path     string            `json:"path"`
	Title    string            `json:"title"`
	Children []AiAgentMenuItem `json:"children"`
}

type AiAgentHistoryMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AiAgentPlanRequest struct {
	Message      string                  `json:"message"`
	Menus        []AiAgentMenuItem       `json:"menus"`
	Tools        []AiAgentTool           `json:"tools"`
	Attachments  []AiAgentAttachment     `json:"attachments"`
	History      []AiAgentHistoryMessage `json:"history"`
	Instructions string                  `json:"instructions"`
}

type AiAgentToolCall struct {
	Tool      string         `json:"tool"`
	Arguments map[string]any `json:"arguments"`
	Message   string         `json:"message"`
}

type AiAgentPlanResponse struct {
	Message     string            `json:"message"`
	ToolCalls   []AiAgentToolCall `json:"toolCalls"`
	DoneMessage string            `json:"doneMessage"`
}

type openAIChatRequest struct {
	Model          string              `json:"model"`
	Messages       []openAIChatMessage `json:"messages"`
	Temperature    float64             `json:"temperature"`
	ResponseFormat map[string]string   `json:"response_format"`
}

type openAIChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

func PlanAiAgent(request AiAgentPlanRequest) (AiAgentPlanResponse, error) {
	if len(request.Tools) == 0 {
		return AiAgentPlanResponse{}, errors.New("可用工具为空，无法规划操作")
	}

	apiKey := firstEnv("OPENAI_API_KEY", "OPENAI_TOKEN")
	if apiKey == "" {
		return AiAgentPlanResponse{}, errors.New("未配置 OPENAI_API_KEY 或 OPENAI_TOKEN")
	}

	model := firstEnv("OPENAI_MODEL", "LLM_MODEL")
	if model == "" {
		model = "gpt-4o-mini"
	}

	baseURL := firstEnv("OPENAI_BASE_URL", "OPENAI_URI")
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}

	content, err := requestOpenAIPlan(baseURL, apiKey, model, request)
	if err != nil {
		return AiAgentPlanResponse{}, err
	}

	plan, err := parsePlanResponse(content)
	if err != nil {
		return AiAgentPlanResponse{}, err
	}

	if len(plan.ToolCalls) == 0 {
		return AiAgentPlanResponse{}, errors.New("LLM 没有返回可执行工具")
	}

	if err := validateToolCalls(plan.ToolCalls, request.Tools); err != nil {
		return AiAgentPlanResponse{}, err
	}

	return plan, nil
}

func requestOpenAIPlan(baseURL string, apiKey string, model string, request AiAgentPlanRequest) (string, error) {
	payload := openAIChatRequest{
		Model:       model,
		Temperature: 0.2,
		ResponseFormat: map[string]string{
			"type": "json_object",
		},
		Messages: []openAIChatMessage{
			{
				Role:    "system",
				Content: buildSystemPrompt(),
			},
			{
				Role:    "user",
				Content: buildUserPrompt(request),
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := strings.TrimRight(baseURL, "/") + "/chat/completions"
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("LLM 请求失败：%s，%s", resp.Status, string(respBody))
	}

	var chatResp openAIChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return "", err
	}

	if chatResp.Error != nil {
		return "", errors.New(chatResp.Error.Message)
	}

	if len(chatResp.Choices) == 0 {
		return "", errors.New("LLM 响应为空")
	}

	return chatResp.Choices[0].Message.Content, nil
}

func buildSystemPrompt() string {
	return `你是一个 Web UI Agent planner。你只负责返回工具调用计划，不直接操作数据库。

硬性规则：
1. 只能使用用户请求中 tools 数组提供的 tool id，不要编造工具。
2. 如果用户要求创建、新增计划，优先先调用 menu.myPlans，再调用 todo.create。
3. 如果用户上传论文/文档并要求拆分实现计划，先调用 menu.myPlans，再调用 todo.batchCreate。
4. todo.batchCreate.arguments.plans 必须是数组。每个 plan 必须包含 userName、title、content、planStartDate、planFinishDate、planTag、items。
5. 每个 item 必须包含 title、content、remark、startDate、endDate。
6. 日期格式必须是 YYYY-MM-DD。
7. 只返回 JSON，不要 Markdown，不要解释。

返回 JSON 格式：
{
  "message": "简短说明",
  "toolCalls": [
    { "tool": "工具 id", "arguments": {}, "message": "步骤提示" }
  ],
  "doneMessage": "完成提示"
}`
}

func buildUserPrompt(request AiAgentPlanRequest) string {
	type promptPayload struct {
		Message      string                  `json:"message"`
		Menus        []AiAgentMenuItem       `json:"menus"`
		Tools        []AiAgentTool           `json:"tools"`
		Attachments  []AiAgentAttachment     `json:"attachments"`
		History      []AiAgentHistoryMessage `json:"history"`
		Instructions string                  `json:"instructions"`
		Today        string                  `json:"today"`
	}

	payload := promptPayload{
		Message:      request.Message,
		Menus:        request.Menus,
		Tools:        request.Tools,
		Attachments:  trimAttachments(request.Attachments),
		History:      request.History,
		Instructions: request.Instructions,
		Today:        time.Now().Format("2006-01-02"),
	}

	data, _ := json.Marshal(payload)
	return string(data)
}

func trimAttachments(attachments []AiAgentAttachment) []AiAgentAttachment {
	const maxContentLength = 12000
	result := make([]AiAgentAttachment, 0, len(attachments))

	for _, item := range attachments {
		if item.Encoding == "text" && len(item.Content) > maxContentLength {
			item.Content = item.Content[:maxContentLength]
		}
		result = append(result, item)
	}

	return result
}

func parsePlanResponse(content string) (AiAgentPlanResponse, error) {
	var plan AiAgentPlanResponse
	if err := json.Unmarshal([]byte(content), &plan); err == nil {
		return plan, nil
	}

	start := strings.Index(content, "{")
	end := strings.LastIndex(content, "}")
	if start < 0 || end <= start {
		return AiAgentPlanResponse{}, errors.New("LLM 未返回合法 JSON")
	}

	if err := json.Unmarshal([]byte(content[start:end+1]), &plan); err != nil {
		return AiAgentPlanResponse{}, err
	}

	return plan, nil
}

func validateToolCalls(calls []AiAgentToolCall, tools []AiAgentTool) error {
	allowed := map[string]bool{}
	dangerous := map[string]bool{}

	for _, tool := range tools {
		allowed[tool.Id] = true
		if tool.Dangerous {
			dangerous[tool.Id] = true
		}
	}

	for _, call := range calls {
		if !allowed[call.Tool] {
			return fmt.Errorf("LLM 返回了未注册工具：%s", call.Tool)
		}
		if dangerous[call.Tool] {
			return fmt.Errorf("危险工具需要人工确认：%s", call.Tool)
		}
	}

	return nil
}

func firstEnv(keys ...string) string {
	for _, key := range keys {
		value := strings.TrimSpace(os.Getenv(key))
		if value != "" {
			return value
		}
	}

	return ""
}
