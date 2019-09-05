package main

import (
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

// 2.3 間の要素を削除: 単方向連結リストにおいて、間の要素（必ずしもちょうど中央というわけではなく、最初と最後の要素以外）
// で、その要素のみアクセス可能であるとします。その要素を削除するアルゴリズムを実装してください
// 
// 例
// 入力: a -> b -> c -> d -> e -> f という連結リストのcが与えられます
// 出力: 何も返しませんが、リストはa -> b -> d -> e -> fのように見えます
func main() {}

func deleteElm(elm *list.Node) {
	if (elm.Next == nil) {
		// 末尾の場合は削除処理はできない
		return
	}
	next := elm.Next
	elm.Next = next.Next
	elm.Value = next.Value
}
