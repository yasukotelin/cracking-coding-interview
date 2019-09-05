package main

import (
	"testing"
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestPartition(t *testing.T) {

	x := 5

	sigList := list.NewSingleList()
	sigList.Add(3)
	sigList.Add(5)
	sigList.Add(8)
	sigList.Add(5)
	sigList.Add(10)
	sigList.Add(2)
	sigList.Add(1)

	exp := list.NewSingleList()
	exp.Add(3)
	exp.Add(2)
	exp.Add(1)
	exp.Add(5)
	exp.Add(8)
	exp.Add(5)
	exp.Add(10)

	act := partition(sigList, x)

	expStr, _ := exp.String()
	actStr, _ := act.String()

	if expStr != actStr {
		t.Errorf("expected is %+v but actual is %+v", expStr, actStr)
	}
}
