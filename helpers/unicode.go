package helpers

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func NormalizeUnicode(str string) string {
	var result string
	for _, c := range str {
		result += string(unicode.ToLower(c))
	}
	return norm.NFKD.String(result)
}
