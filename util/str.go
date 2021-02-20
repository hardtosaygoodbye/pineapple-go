package util

import (
	"errors"
	"strings"
)

// SplitContent 截取某一位置的文本
func SplitContent(content string, preContent string, nextContent string) (result string, err error) {
	if !strings.Contains(content, preContent) || !strings.Contains(content, nextContent) {
		err = errors.New("该位置描述错误")
		return
	}
	result = strings.Split(content, preContent)[1]
	result = strings.Split(result, nextContent)[0]
	return
}
