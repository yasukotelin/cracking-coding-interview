package main

import (
	"reflect"
	"testing"

	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestAdd(t *testing.T) {

	// 617
	l1 := list.NewSingleList()
	l1.Add(7)
	l1.Add(1)
	l1.Add(6)

	// 295
	l2 := list.NewSingleList()
	l2.Add(5)
	l2.Add(9)
	l2.Add(2)

	// 617 + 219 = 912
	exp := list.NewSingleList()
	exp.Add(2)
	exp.Add(1)
	exp.Add(9)

	act := add(l1, l2)

	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("expected is %+v but actual is %+v", exp, act)
	}
}

func TestFmtToNumber1(t *testing.T) {
	// 617
	l1 := list.NewSingleList()
	l1.Add(7)
	l1.Add(1)
	l1.Add(6)

	exp := 617
	act := fmtToNumber1(l1)

	if exp != act {
		t.Fatalf("expected is %v but actual is %v", exp, act)
	}
}

func TestFmtToNumber2(t *testing.T) {
	// 617
	l1 := list.NewSingleList()
	l1.Add(6)
	l1.Add(1)
	l1.Add(7)

	exp := 617
	act := fmtToNumber2(l1)

	if exp != act {
		t.Fatalf("expected is %v but actual is %v", exp, act)
	}
}
