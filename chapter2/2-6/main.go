package main

import (
	"reflect"

	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

/**
2.6 回文: 連結リストが回文（先頭から巡回しても末尾から巡回しても、各ノードの要素が
まったく同じになっている）かどうかを調べる関数を実感してください
**/
func main() {}

func isParindrome(sigList *list.SingleList) bool {
	reversed := reverse(sigList)
	return reflect.DeepEqual(sigList, reversed)
}

func reverse(sigList *list.SingleList) *list.SingleList {
	reversed := list.NewSingleList()

	sigNode := sigList.Head
	revNode := list.NewNode(sigNode.Value)
	reversed.Tail = revNode

	sigNode = sigNode.Next

	for sigNode != nil {
		newNode := list.NewNode(sigNode.Value)
		newNode.Next = revNode

		revNode = newNode
		sigNode = sigNode.Next
	}

	reversed.Head = revNode

	return reversed
}
