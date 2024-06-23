/*
Create: 2023/2/19
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package main

import (
	"common"
	"github.com/gin-gonic/gin"
)

// 注册路由

var routes = []common.Route{
	{
		Path:   "/",
		Des:    "",
		Method: "",
	},
}

func Router(s *gin.Engine) {
	group := s.Group(common.API(API), common.Load(Module.enable))
	{
		group.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"routes": routes})
		})
	}
}
