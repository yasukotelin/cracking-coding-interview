// main 1.9
// 片方の文字列が、もう片方の文字列の一部分になっているかどうかを調べるメソッド「isSubstring」
// が使えると仮定します。2つの文字列s1とs2が与えられたとき、isSubstringメソッドを一度だけ使って
// s2がs1を回転させたものかどうかを判定するコードを書いてください。
// （例えば、「waterbottle」は「erbottlewat」を回転させたものになっています）
package main

import (
	"strings"
)

func main() {}

// isCircledString はs2がs1を回転させたものかどうかを返却する
func isCircledString(s1, s2 string) bool {
	// 回転している場合、s2 = s1s1の部分文字列になっている
	return isSubString(s1 + s1, s2)
}

func isSubString(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}
