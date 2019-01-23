package utils

import (
    "net/url"
    "strings"
)

// EncodeURIComponent 编码，实现上类似于js的
func EncodeURIComponent(str string) string {
	u := url.QueryEscape(str)
	// 处理空格
    u = strings.Replace(u, "+", "%20", -1)
    return u
}
