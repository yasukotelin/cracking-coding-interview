package main

import (
	"strconv"
	"strings"
)

func main() {
}

func simpleCompression(s string) string {
	builder := new(strings.Builder)

	runes := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		cur := runes[i]

		if i == len(s)-1 {
			builder.WriteRune(cur)
			builder.WriteString(strconv.Itoa(count))
			break
		}

		next := runes[i+1]
		if cur != next {
			builder.WriteRune(cur)
			builder.WriteString(strconv.Itoa(count))
			count = 0
		}
	}

	result := builder.String()
	if len(result) >= len(s) {
		return s
	} else {
		return builder.String()
	}
}
