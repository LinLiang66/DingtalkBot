package utils

import (
	"regexp"
	"strings"
)

func CutPrefix(s, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return strings.TrimPrefix(s, prefix), true
	}
	return s, false
}

func EitherCutPrefix(s string, prefix ...string) (string, bool) {

	// 任一前缀匹配则返回剩余部分
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return strings.TrimPrefix(s, p), true
		}
	}
	return s, false
}

// TrimEqual trim space and equal
func TrimEqual(s, prefix string) (string, bool) {
	if strings.TrimSpace(s) == prefix {
		return "", true
	}
	return s, false
}

func EitherTrimEqual(s string, prefix ...string) (string, bool) {
	// 任一前缀匹配则返回剩余部分
	for _, p := range prefix {
		if strings.TrimSpace(s) == p {
			return "", true
		}
	}
	return s, false
}

// ContainsSpecificContent 使用正则表达式检查文本内容是否包含特定内容
func ContainsSpecificContent(text string, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(text), nil
}

// ContainsSpecificContentV2 使用正则表达式检查文本内容是否包含特定内容
func ContainsSpecificContentV2(text string, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func isEmpty(str *string) bool {
	return str == nil || *str == ""
}
