package main

import (
	"fmt"
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

// 2.2 後ろからK番目を返す
// 単方向連結リストにおいて、末尾から数えてk番目の要素を見つける
// アルゴリズムを自走してください
func main() {
}

func findFromBehind(singleList *list.SingleList, k int) (int, error){
	// 1. singleListのカウント取得
	// 2. 後ろからk番目が先頭から何番目かを取得
	// 3. 対象まで再度走査。要素を返却する

	length := Length(singleList)
	targetPos := length - k

	var cnt int
	node := singleList.Head
	for node != nil {
		if (cnt == targetPos) {
			return node.Value, nil
		}
		cnt++
		node = node.Next
	}

	return 0, fmt.Errorf("k(%d) is over the node length", k)
}

func Length(singleList *list.SingleList) int {
	var length int
	node := singleList.Head
	for node != nil {
		length++
		node = node.Next
	}
	return length
}
