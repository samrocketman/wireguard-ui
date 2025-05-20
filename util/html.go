package util

import (
	"html"
)

func EscapeHtmlCode(s string) string {
	encodedString := html.EscapeString(s)
	return encodedString
}
