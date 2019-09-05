package main

import (
	"strconv"
	"strings"

	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

// 2.5 リストで表された2数の和: 各ノードの要素が1桁の数である連結リストで表された2つの数があります。
// 一の位がリストの先頭になるように、各位の数は逆順に並んでいます。
// このとき2つの数の和を求め、それを連結リストで表したものを返す関数を書いてください。
//
// 例
// 入力: (7 -> 1 -> 6) + (5 -> 9 -> 2) -> 617 + 295
// 出力: (2 -> 1 -> 9) -> 912
//
// 発展問題
// 上位の桁から順方向に連結されたリストを用いて、同様に解いてください
// 入力: (6 -> 1 -> 7) + (2 -> 9 -> 5) -> 617 + 295
// 出力: (9 -> 1 -> 2) -> 912
func main() {}

func add(l1, l2 *list.SingleList) *list.SingleList {
	result := list.NewSingleList()

	l1Num := fmtToNumber1(l1)
	l2Num := fmtToNumber1(l2)

	resNum := l1Num + l2Num

	resNumRunes := []rune(strconv.Itoa(resNum))

	for i := len(resNumRunes) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(resNumRunes[i]))
		result.Add(num)
	}

	return result
}

func fmtToNumber1(l *list.SingleList) int {
	var s string
	node := l.Head
	for node != nil {
		s = strconv.Itoa(node.Value) + s
		node = node.Next
	}
	num, _ := strconv.Atoi(s)

	return num
}

func fmtToNumber2(l *list.SingleList) int {
	var builder strings.Builder
	node := l.Head
	for node != nil {
		builder.WriteString(strconv.Itoa(node.Value))
		node = node.Next
	}
	numStr := builder.String()
	num, _ := strconv.Atoi(numStr)

	return num
}
