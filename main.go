package main

import (
	"net/http"
	"toGO/common"
	"toGO/config"
	"toGO/contr"
	"toGO/contr/req"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(config.GetCorsConfig())

	menuController := contr.GetMenuController()

	planTodoController := contr.GetPlanTodoController()

	app.POST("/menu/page_list", func(c *gin.Context) {
		params := req.MenuPageRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be MenuPageRequest"))
		} else {
			list := menuController.PageList(params)
			c.JSON(http.StatusOK, list)
		}

	})

	app.POST("/menu/list", func(c *gin.Context) {
		params := req.MenuPageRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be MenuPageRequest"))
		} else {
			list := menuController.PageList(params)
			c.JSON(http.StatusOK, list)
		}
	})

	app.GET("/menu/list", func(c *gin.Context) {
		name := c.Query("name")
		resp := menuController.List(name)
		c.JSON(http.StatusOK, resp)
	})

	app.GET("/dashboard/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(dashboardHomeData()))
	})

	app.GET("/todo/overview_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(todoOverviewData()))
	})

	app.GET("/todo/board_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(todoBoardData()))
	})

	app.GET("/todo/calendar_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(todoCalendarData()))
	})

	app.GET("/todo/review_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(todoReviewData()))
	})

	app.GET("/learning/path_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(learningPathData()))
	})

	app.GET("/learning/schedule_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(learningScheduleData()))
	})

	app.GET("/learning/materials_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(learningMaterialsData()))
	})

	app.GET("/learning/review_data", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.Success(learningReviewData()))
	})

	app.POST("/plan/page_list", func(c *gin.Context) {
		params := req.TodoPageRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be TodoPageRequest"))
		} else {
			c.JSON(http.StatusOK, planTodoController.PageList(params))
		}
	})

	app.POST("/plan/create", func(c *gin.Context) {
		params := req.PlanCreateRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be PlanCreateRequest"))
		} else {
			c.JSON(http.StatusOK, planTodoController.CreatePlan(params))
		}
	})

	_ = app.Run(":8000")

}
