/*
Create: 2023/2/18
Project: Atlas
Github: https://github.com/landers1037
Copyright Renj
*/

package common

// 生成extra信息

func Extra(api string, des string) map[string]interface{} {
	return map[string]interface{}{
		"api": API(api),
		"des": des,
	}
}
