package handlers

import (
	"regexp"
	"strconv"
	"strings"
)

func processNewLine(msg string) string {
	return strings.Replace(msg, "\\n", `
`, -1)
}

func processQuote(msg string) string {
	return strings.Replace(msg, "\\\"", "\"", -1)
}

// 将字符中 \u003c 替换为 <  等等
func processUnicode(msg string) string {
	regex := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	return regex.ReplaceAllStringFunc(msg, func(s string) string {
		r, _ := regexp.Compile(`\\u`)
		s = r.ReplaceAllString(s, "")
		i, _ := strconv.ParseInt(s, 16, 32)
		return string(rune(i))
	})
}

func cleanTextBlock(msg string) string {
	msg = processNewLine(msg)
	msg = processUnicode(msg)
	msg = processQuote(msg)
	return msg
}
