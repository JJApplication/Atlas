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
	Name = "ModuleInflux"
	API  = "influx"
	DES  = "InfluxDB时序数据查询模块"
)

type moduleSys struct {
	enable bool
}

func (m *moduleSys) Name() string {
	return Name
}

func (m *moduleSys) Hooks(s interface{}) {
	Router(s.(*gin.Engine))
}

func (m *moduleSys) Enable() {
	m.enable = true
}

func (m *moduleSys) Disable() {
	m.enable = false
}

func (m *moduleSys) Extra() map[string]interface{} {
	return common.Extra(API, DES)
}

var Module = moduleSys{}
