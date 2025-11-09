package main

import (
	"net/http"
	"toGO/common"
	"toGO/contr/req"

	"github.com/gin-gonic/gin"
)
import "toGO/contr"

func main() {
	app := gin.Default()

	menuController := contr.GetMenuController()

	app.POST("/menu/page_list", func(c *gin.Context) {
		params := req.MenuPageRequest{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusOK, common.Error("the body should be MenuPageRequest"))
		} else {
			list := menuController.PageList(params)
			c.JSON(http.StatusOK, list)
		}

	})

	_ = app.Run(":8000")

}
