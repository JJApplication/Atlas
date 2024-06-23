/*
Create: 2023/2/17
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package main

import (
	"common"
	"github.com/gin-gonic/gin"
)

const (
	Name = "ModuleNginx"
	API  = "nginx"
	DES  = "Nginx配置操作模块"
)

type module struct {
	enable bool
}

func (m *module) Name() string {
	return Name
}

func (m *module) Hooks(s interface{}) {
	Router(s.(*gin.Engine))
}

func (m *module) Enable() {
	m.enable = true
}

func (m *module) Disable() {
	m.enable = false
}

func (m *module) Extra() map[string]interface{} {
	return common.Extra(API, DES)
}

var Module = module{}
