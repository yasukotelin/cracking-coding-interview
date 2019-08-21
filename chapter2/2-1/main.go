package main

import (
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func main() {}

func deleteDuplicate(sigList *list.SingleList) {
	// 重複要素の探索
	duplValues := make(map[int]int)

	var prevNode *list.Node = nil
	node := sigList.Head

	for node != nil {
		if _, ok := duplValues[node.Value]; ok {
			// 重複しているので削除処理
			prevNode.Next = node.Next
			// fmt.Printf("%+v\n", prevNode)
			node = node.Next
		} else {
			duplValues[node.Value] = 1
			prevNode = node
			node = node.Next
		}
	}

	sigList.Tail = prevNode
}
