package main

import (
	"encoding/base64"
	"regexp"
	"unicode/utf8"
)

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

//判断是否是emoji表情
func JudgeEmoji(content string) bool {
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size == 4 {
			return true
		}
	}
	return false
}

//不能判断一定是，可以判断一定不是
func JudgeBase64(str string) bool {
	pattern := "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	if !(len(str)%4 == 0 && matched) {
		return false
	}
	unCodeStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return false
	}
	tranStr := base64.StdEncoding.EncodeToString(unCodeStr)
	//return str==base64.StdEncoding.EncodeToString(unCodeStr)
	if str == tranStr {
		return true
	}
	return false
}