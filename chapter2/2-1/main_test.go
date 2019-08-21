package main

import (
	"testing"

	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestDeleteDuplicate1(t *testing.T) {
	// arrange
	sigList := list.NewSingleList()
	sigList.Add(1)
	sigList.Add(33)
	sigList.Add(21)
	sigList.Add(33)
	deleteDuplicate(sigList)

	exp := list.NewSingleList()
	exp.Add(1)
	exp.Add(33)
	exp.Add(21)

	sigListStr, _ := sigList.String()
	expStr, _ := exp.String()

	if sigListStr != expStr {
		t.Fatalf("expected is %#v but actual is %#v", expStr, sigListStr)
	}
}

func TestDeleteDuplicate2(t *testing.T) {
	// arrange
	sigList := list.NewSingleList()
	sigList.Add(12)
	sigList.Add(3)
	sigList.Add(1)
	sigList.Add(4)
	sigList.Add(12)
	sigList.Add(31)
	sigList.Add(11)
	deleteDuplicate(sigList)

	exp := list.NewSingleList()
	exp.Add(12)
	exp.Add(3)
	exp.Add(1)
	exp.Add(4)
	exp.Add(31)
	exp.Add(11)

	sigListStr, _ := sigList.String()
	expStr, _ := exp.String()

	if sigListStr != expStr {
		t.Fatalf("expected is %#v but actual is %#v", expStr, sigListStr)
	}
}
