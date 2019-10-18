package simpleparser

import (
	"strings"
)

// SafeHTML return html
func SafeHTML(str string) string {
	sp := strings.Split(str, "\n")
	var HTML string
	for _, val := range sp {
		HTML= HTML+`<p>`+val+`</p>`

	}
	return HTML
}