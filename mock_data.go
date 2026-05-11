package main

func dashboardHomeData() any {
	return map[string]any{
		"metrics": []map[string]any{
			{"label": "进行中计划", "value": "12", "trend": "+3", "hint": "较上周新增", "type": "success"},
			{"label": "待处理待办", "value": "36", "trend": "8 项临期", "hint": "建议今天完成", "type": "warning"},
			{"label": "本周完成率", "value": "68%", "trend": "+12%", "hint": "节奏保持稳定", "type": "primary"},
			{"label": "AI 辅助次数", "value": "24", "trend": "可提升", "hint": "可继续补充自动化动作", "type": "info"},
		},
		"weeklyPlans": []map[string]any{
			{"title": "完善计划创建流程", "owner": "产品", "date": "今天", "progress": 82, "status": "is-active"},
			{"title": "补齐待办筛选和状态", "owner": "前端", "date": "明天", "progress": 55, "status": "is-normal"},
			{"title": "设计计划复盘看板", "owner": "运营", "date": "周五", "progress": 36, "status": "is-muted"},
		},
		"features": []map[string]any{
			{"title": "计划状态流转", "desc": "待开始、进行中、已完成、延期统一管理。", "icon": "Operation"},
			{"title": "智能拆解待办", "desc": "用 AI 根据计划内容生成可执行任务。", "icon": "ChatDotRound"},
			{"title": "进度统计看板", "desc": "按用户、标签、时间维度观察完成率。", "icon": "DataAnalysis"},
			{"title": "到期提醒", "desc": "临期任务突出展示，并支持快捷筛选。", "icon": "Timer"},
		},
	}
}

func todoOverviewData() any {
	return map[string]any{
		"metrics": []map[string]any{
			{"label": "计划总数", "value": "42", "percent": 84},
			{"label": "按期推进", "value": "31", "percent": 74},
			{"label": "临期待办", "value": "8", "percent": 38},
			{"label": "延期事项", "value": "3", "percent": 14},
		},
		"risks": []map[string]any{
			{"title": "移动端适配复核", "desc": "明天到期，缺少验收截图", "level": "临期", "type": "warning"},
			{"title": "创建计划接口联调", "desc": "后端字段仍有变更风险", "level": "阻塞", "type": "danger"},
			{"title": "AI 自动拆解能力", "desc": "需要补充提示词和工具协议", "level": "关注", "type": "info"},
			{"title": "复盘报表口径", "desc": "完成率定义需要产品确认", "level": "确认", "type": "primary"},
		},
		"owners": []map[string]any{
			{"name": "产品", "count": 12, "focus": "需求拆解", "load": 82},
			{"name": "前端", "count": 16, "focus": "页面交互", "load": 76},
			{"name": "后端", "count": 9, "focus": "接口联调", "load": 58},
			{"name": "运营", "count": 6, "focus": "复盘分析", "load": 42},
		},
	}
}

func todoBoardData() any {
	return []map[string]any{
		{"key": "todo", "title": "待开始", "label": "Backlog", "type": "info", "items": []map[string]any{
			{"title": "配置计划标签体系", "desc": "定义标签颜色、筛选和统计口径。", "owner": "产品", "date": "5月12日", "priority": "P2", "priorityType": "info"},
			{"title": "补充批量导入入口", "desc": "支持从表格粘贴或上传导入待办。", "owner": "前端", "date": "5月13日", "priority": "P3", "priorityType": "info"},
		}},
		{"key": "doing", "title": "进行中", "label": "Doing", "type": "primary", "items": []map[string]any{
			{"title": "完善创建计划弹窗", "desc": "补齐校验、待办列表和提交反馈。", "owner": "前端", "date": "今天", "priority": "P1", "priorityType": "danger"},
			{"title": "联调分页查询接口", "desc": "统一分页字段和错误提示。", "owner": "后端", "date": "今天", "priority": "P1", "priorityType": "danger"},
		}},
		{"key": "review", "title": "验收中", "label": "Review", "type": "warning", "items": []map[string]any{
			{"title": "移动端布局检查", "desc": "重点检查菜单、表格和弹窗宽度。", "owner": "测试", "date": "明天", "priority": "P2", "priorityType": "warning"},
		}},
		{"key": "done", "title": "已完成", "label": "Done", "type": "success", "items": []map[string]any{
			{"title": "工作台视觉重构", "desc": "首页指标、推进列表和功能入口已完成。", "owner": "前端", "date": "昨天", "priority": "P2", "priorityType": "success"},
		}},
	}
}

func todoCalendarData() any {
	days := make([]map[string]any, 0, 35)
	taskMap := map[int][]string{
		3:  []string{"需求确认"},
		6:  []string{"接口联调", "视觉验收"},
		10: []string{"上线准备"},
		13: []string{"风险复盘"},
		18: []string{"看板迭代"},
		22: []string{"AI 拆解"},
		27: []string{"数据统计"},
	}
	for i := 1; i <= 35; i++ {
		tasks := taskMap[i]
		if tasks == nil {
			tasks = []string{}
		}
		days = append(days, map[string]any{"date": i, "today": i == 10, "tasks": tasks})
	}
	return map[string]any{
		"days": days,
		"agenda": []map[string]any{
			{"time": "09:30", "title": "确认计划状态字段", "owner": "产品 / 后端"},
			{"time": "14:00", "title": "列表页联调回归", "owner": "前端 / 测试"},
			{"time": "17:30", "title": "整理明日优先级", "owner": "团队"},
		},
	}
}

func todoReviewData() any {
	return map[string]any{
		"score":     86,
		"scoreDesc": "计划拆解清晰，延期集中在接口联调阶段。",
		"reasons": []map[string]any{
			{"name": "需求边界变更", "value": 72},
			{"name": "接口联调延迟", "value": 64},
			{"name": "验收标准不清", "value": 46},
			{"name": "资源排期冲突", "value": 34},
		},
		"records": []map[string]any{
			{"plan": "计划列表改版", "result": "完成", "reason": "按预期推进", "action": "抽出通用页面结构用于后续模块", "owner": "前端"},
			{"plan": "创建计划流程", "result": "部分延期", "reason": "后端接口字段未稳定", "action": "先固定前端模型，保留适配层", "owner": "后端"},
			{"plan": "AI 操作助手", "result": "待验证", "reason": "缺少真实 Agent 服务", "action": "补充工具协议和失败兜底提示", "owner": "产品"},
		},
	}
}

func learningScheduleData() any {
	return []map[string]any{
		{"day": "周一", "total": "95 分钟", "items": []map[string]any{
			{"title": "Vue 响应式复盘", "time": "45 分钟", "mode": "笔记", "typeLabel": "复盘", "type": "primary"},
			{"title": "组件通信练习", "time": "50 分钟", "mode": "项目", "typeLabel": "产出", "type": "warning"},
		}},
		{"day": "周二", "total": "60 分钟", "items": []map[string]any{
			{"title": "双指针专项", "time": "60 分钟", "mode": "刷题", "typeLabel": "练习", "type": "success"},
		}},
		{"day": "周三", "total": "40 分钟", "items": []map[string]any{
			{"title": "英文技术文章", "time": "40 分钟", "mode": "阅读", "typeLabel": "输入", "type": "info"},
		}},
	}
}

func learningMaterialsData() any {
	return []map[string]any{
		{"title": "Vue 3 官方指南", "desc": "用于复习响应式、组件和组合式 API 的核心概念。", "topic": "前端工程", "progress": "已读 70%", "status": "在学", "sourceType": "文档", "url": "https://vuejs.org/guide/introduction.html", "files": []string{}, "type": "primary", "icon": "Document"},
		{"title": "TypeScript Handbook", "desc": "补齐泛型、联合类型和工具类型的建模能力。", "topic": "类型系统", "progress": "已读 45%", "status": "在学", "sourceType": "文档", "url": "https://www.typescriptlang.org/docs/", "files": []string{}, "type": "warning", "icon": "Collection"},
		{"title": "算法题单", "desc": "按数组、栈队列、动态规划等主题归档练习。", "topic": "算法训练", "progress": "完成 18 题", "status": "练习", "sourceType": "题单", "url": "", "files": []string{"双指针错题.md"}, "type": "success", "icon": "Link"},
		{"title": "英文技术分享", "desc": "用作听力、表达和技术词汇积累材料。", "topic": "英语输入", "progress": "本周 2 次", "status": "积累", "sourceType": "视频", "url": "", "files": []string{}, "type": "info", "icon": "VideoPlay"},
	}
}

func learningReviewData() any {
	return map[string]any{
		"reviewItems": []map[string]any{
			{"title": "computed 与 watch 的边界", "desc": "需要补一个实际项目例子验证选择依据。", "level": "概念", "type": "warning"},
			{"title": "双指针窗口收缩条件", "desc": "错题集中同类问题重复出现。", "level": "练习", "type": "danger"},
			{"title": "TypeScript 泛型约束", "desc": "整理 3 个常用业务建模场景。", "level": "产出", "type": "primary"},
		},
		"mastery": []map[string]any{
			{"label": "Vue 3 基础", "hint": "概念稳定，可进入项目练习", "value": 82},
			{"label": "TypeScript", "hint": "需要继续强化类型推导", "value": 56},
			{"label": "算法训练", "hint": "题型识别还不稳定", "value": 48},
			{"label": "英语输入", "hint": "保持低强度连续积累", "value": 64},
		},
	}
}

func learningPathData() any {
	return map[string]any{
		"tracks": map[string]any{
			"frontend": map[string]any{"label": "前端工程", "stages": []map[string]any{
				{"order": "01", "title": "Vue 3 与组合式 API", "goal": "补齐响应式、组件通信和页面状态组织方式。", "duration": "5 天", "output": "完成 2 个可复用组件", "progress": 86, "status": "推进中", "type": "primary"},
				{"order": "02", "title": "TypeScript 建模", "goal": "把计划、条目、复盘记录抽象成稳定类型。", "duration": "4 天", "output": "整理实体与接口定义", "progress": 58, "status": "待强化", "type": "warning"},
				{"order": "03", "title": "工程化与质量检查", "goal": "熟悉 Vite、ESLint、构建和发布前检查。", "duration": "3 天", "output": "形成个人检查清单", "progress": 32, "status": "待开始", "type": "info"},
			}},
			"algorithm": map[string]any{"label": "算法训练", "stages": []map[string]any{
				{"order": "01", "title": "数组与双指针", "goal": "建立题型识别和边界处理习惯。", "duration": "3 天", "output": "完成 12 道题与错题记录", "progress": 72, "status": "推进中", "type": "primary"},
				{"order": "02", "title": "栈、队列与单调结构", "goal": "掌握状态维护和入栈出栈条件。", "duration": "4 天", "output": "整理模板与变体", "progress": 44, "status": "待强化", "type": "warning"},
				{"order": "03", "title": "动态规划入门", "goal": "练习状态定义、转移方程和初始化。", "duration": "6 天", "output": "输出 1 篇复盘笔记", "progress": 18, "status": "待开始", "type": "info"},
			}},
			"english": map[string]any{"label": "英语输入", "stages": []map[string]any{
				{"order": "01", "title": "技术文章精读", "goal": "围绕前端工程主题积累表达和术语。", "duration": "每周 3 次", "output": "摘录 30 个表达", "progress": 65, "status": "推进中", "type": "primary"},
				{"order": "02", "title": "听力影子跟读", "goal": "提升技术分享和会议语境下的理解速度。", "duration": "每周 2 次", "output": "录制 5 分钟复述", "progress": 40, "status": "待强化", "type": "warning"},
				{"order": "03", "title": "英文周报输出", "goal": "用英文总结本周学习内容和问题。", "duration": "每周 1 次", "output": "提交一份学习周报", "progress": 24, "status": "待开始", "type": "info"},
			}},
		},
		"weekPlan": []map[string]any{
			{"id": 1, "day": "周一", "title": "Vue 组合式 API 复盘", "time": "45 分钟", "mode": "笔记整理", "focus": "输入", "type": "primary"},
			{"id": 2, "day": "周二", "title": "数组双指针专项", "time": "60 分钟", "mode": "刷题训练", "focus": "练习", "type": "success"},
			{"id": 3, "day": "周三", "title": "TypeScript 类型建模", "time": "50 分钟", "mode": "项目实践", "focus": "产出", "type": "warning"},
			{"id": 4, "day": "周五", "title": "英文技术文章精读", "time": "40 分钟", "mode": "阅读摘录", "focus": "积累", "type": "info"},
		},
		"dependencies": []map[string]any{
			{"title": "先补齐实体字段", "desc": "学习路径需要目标、阶段、依赖、复盘时间等字段。", "level": "高", "type": "danger"},
			{"title": "区分学习条目与普通待办", "desc": "学习条目需要记录掌握度、复习次数和材料链接。", "level": "中", "type": "warning"},
			{"title": "建立复盘口径", "desc": "完成不等于掌握，需要增加回忆和实践验证。", "level": "中", "type": "primary"},
		},
		"reviews": []map[string]any{
			{"title": "Vue 响应式边界", "desc": "对 ref、reactive、computed 的适用场景做一次对照。", "retention": 76},
			{"title": "双指针错题集", "desc": "重点回看窗口收缩条件和重复元素处理。", "retention": 58},
			{"title": "英文技术词汇", "desc": "把本周摘录放进下一轮间隔复习。", "retention": 42},
		},
	}
}
