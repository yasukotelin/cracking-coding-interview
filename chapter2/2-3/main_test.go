package main

import (
	"testing"
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestDelete1(t *testing.T) {
	elm := list.NewNode(3)

	sigList := list.NewSingleList()
	sigList.Add(1)
	sigList.Add(2)
	sigList.AddNode(elm)
	sigList.Add(4)
	sigList.Add(5)

	expSigList := list.NewSingleList()
	expSigList.Add(1)
	expSigList.Add(2)
	expSigList.Add(4)
	expSigList.Add(5)

	deleteElm(elm)

	sigListStr, _ := sigList.String()
	expSigListStr, _ := expSigList.String()

	if sigListStr != expSigListStr {
		t.Fatalf("expected is %+v but actual is %+v", expSigListStr, sigListStr)
	}
}
