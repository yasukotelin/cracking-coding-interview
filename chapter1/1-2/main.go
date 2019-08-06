package main

import (
	"fmt"
	"sort"
)

func main() {
	var r rune = '世'
	fmt.Println(r)
}

// isSame はs2とs2が完全に同じ文字を使って作られた文字列であるかを判定する
// 例） s1 -> hello s2 -> oelhl => true
//
// ただし、大文字小文字は区別をし、空白は意味のある文字とする
func isSame(s1, s2 string) bool {
	s1Runes := []rune(s1)
	sort.Slice(s1Runes, func(i, j int) bool {
		return s1Runes[i] < s1Runes[j]
	})
	s2Runes := []rune(s2)
	sort.Slice(s2Runes, func(i, j int) bool {
		return s2Runes[i] < s2Runes[j]
	})

	return string(s1Runes) == string(s2Runes)
}
