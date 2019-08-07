package main

import (
	"strings"
)

const whiteSpace = rune(' ')

func main() {
}

// isPalindrome は文字列sが回文の順列であるかを返却する
// 回文とは、文字列が前から読んでも後ろから読んでも同じになる文字のこと
func isPalindrome(s string) bool {

	// 小文字大文字区別なし
	runes := []rune(strings.ToUpper(s))

	hash := make(map[rune]int)
	for _, r := range runes {
		if r == whiteSpace {
			// 空白文字は除外
			continue
		}
		hash[r]++
	}

	// 奇数長の場合奇数の出現は一度だけ可能
	// 偶数の場合必ず偶数のみか奇数が複数回出現するためisUsedOddを同様に使用しても問題ない
	isUsedOdd := false

	for _, cnt := range hash {
		if cnt%2 == 1 && isUsedOdd {
			return false
		} else if cnt%2 == 1 {
			isUsedOdd = true
		}
	}

	return true
}
