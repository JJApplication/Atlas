/*
Create: 2023/2/17
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package common

import (
	"fmt"
	"strings"
)

const (
	Prefix = "/api/module"
)

func API(api string) string {
	if strings.Contains(api, "/") {
		api = strings.TrimLeft(api, "/")
	}
	return fmt.Sprintf("%s/%s", Prefix, api)
}
