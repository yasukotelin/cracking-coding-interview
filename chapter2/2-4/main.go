package main

import (
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

// 2.4 リストの分割: ある数xが与えられたとき、連結リストの要素を並び替え、xより小さいものが前に来るようにするコード
// を書いてください。xがリストに含まれる場合、xの値はxより小さい要素の後にある必要があります（例を参照してください）
// 区切り要素のxは右半分のどこに現れてもかまいません。左半分と右半分のちょうど間にある必要はないということです。
//
// 例
// 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [区切り要素 = 5]
// 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
func main() { }

func partition(sigList *list.SingleList, x int) *list.SingleList {
	smallerList := list.NewSingleList()
	largerList := list.NewSingleList()

	node := sigList.Head
	for node != nil {
		if node.Value < x {
			smallerList.Add(node.Value)
		} else {
			largerList.Add(node.Value)
		}
		node = node.Next
	}

	// merge
	smallerList.AddNode(largerList.Head)
	smallerList.Tail = largerList.Tail

	return smallerList
}
