/*
Create: 2023/2/18
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package common

// 用于返回默认的路由分组
// 告知Apollo接口组后 Apollo通过请求$module/$api获取此分组

type Route struct {
	Path   string // 形如/$api/$path
	Des    string // 描述
	Method string // 请求方式
}
