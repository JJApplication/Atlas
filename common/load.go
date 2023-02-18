/*
Create: 2023/2/17
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package common

import (
	"github.com/gin-gonic/gin"
)

// 是否加载模块

const (
	ModuleDisabled = "ModuleDisabled"
)

func Load(enable bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enable {
			c.JSON(200, gin.H{"status": ModuleDisabled})
			c.Abort()
		}
		c.Next()
	}
}
