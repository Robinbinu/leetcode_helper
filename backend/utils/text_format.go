package utils

import (
	"regexp"
	"strings"
)


func EscapeCodeFieldNewlines(jsonStr string) string {
	re := regexp.MustCompile(`"code":\s*"((?:[^"\\]|\\.|"(?=\s*,))*)"`)
	return re.ReplaceAllStringFunc(jsonStr, func(match string) string {
		// Extract the code string
		prefix := `"code": "`
		start := len(prefix)
		end := len(match) - 1 // remove trailing quote
		code := match[start:end]
		// Replace real newlines with \n and escape quotes
		code = strings.ReplaceAll(code, `\`, `\\`)
		code = strings.ReplaceAll(code, `"`, `\"`)
		code = strings.ReplaceAll(code, "\n", `\n`)
		return prefix + code + `"`
	})
}
