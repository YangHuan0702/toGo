package main

import (
	"net/http"
	"toGO/common"
	"toGO/config"
	"toGO/contr/req"

	"github.com/gin-gonic/gin"
)
import "toGO/contr"

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

	app.GET("/menu/list", func(c *gin.Context) {
		name := c.Param("name")
		resp := menuController.List(name)
		c.JSON(http.StatusOK, resp)
	})

	app.POST("/plan/page_list", func(c *gin.Context) {
		params := req.TodoPageRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be MenuPageRequest"))
		} else {
			c.JSON(http.StatusOK, planTodoController.PageList(params))
		}
	})

	_ = app.Run(":8000")

}
