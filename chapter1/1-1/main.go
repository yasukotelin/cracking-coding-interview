package main

import (
	"sort"
)

func main() {
}

// hasDuplication1 は重複する文字があるかどうかを返却する
func hasDuplication1(s string) bool {
	hash := make(map[rune]int)

	for _, r := range s {
		_, ok := hash[r]
		if ok {
			return true
		}
		hash[r] = 1
	}

	return false
}

// hasDuplication2 は重複する文字があるかどうかを返却する
// ただし、これを実装するのに新たなデータ構造が使えないとする
func hasDuplication2(s string) bool {
	runes := []rune(s)
	for i, _ := range runes {
		for j, _ := range runes {
			if i == j {
				continue
			}
			if runes[i] == runes[j] {
				return true
			}
		}
	}
	return false
}

// hasDuplication3 は重複する文字があるかどうかを返却する
// ただし、これを実装するのに新たなデータ構造が使えないとする
func hasDuplication3(s string) bool {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			return true
		}
	}

	return false
}
