package main

import (
	"reflect"
	"testing"

	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestIsParindrome(t *testing.T) {
	sigList := list.NewSingleList()
	sigList.Add(1)
	sigList.Add(2)
	sigList.Add(3)
	sigList.Add(2)
	sigList.Add(1)

	if !isParindrome(sigList) {
		t.Fatalf("should be true")
	}

	sigList.Add(12)

	if isParindrome(sigList) {
		t.Fatalf("should be false")
	}
}

func TestReverse(t *testing.T) {
	sigList := list.NewSingleList()
	sigList.Add(1)
	sigList.Add(2)
	sigList.Add(3)

	expected := list.NewSingleList()
	expected.Add(3)
	expected.Add(2)
	expected.Add(1)

	reversed := reverse(sigList)

	if !reflect.DeepEqual(expected, reversed) {
		es, _ := expected.String()
		rs, _ := reversed.String()
		t.Fatalf("expected is %v but actual is %v", es, rs)
	}
}
