package main

import (
	"testing"
	"github.com/yasukotelin/cracking-coding-interview/chapter2/list"
)

func TestFindFromBehind1(t *testing.T) {
	singleList := list.NewSingleList()
	singleList.Add(1)
	singleList.Add(33)
	singleList.Add(21)
	singleList.Add(33)
	singleList.Add(108)
	singleList.Add(3)

	cases := []struct {
		k int
		exp int
	} {
		{k: 3, exp: 33},
		{k: 1, exp: 3},
		{k: 5, exp: 33},
	}

	for _, c := range cases {
		act, err := findFromBehind(singleList, c.k)

		if err != nil {
			t.Fatalf("should not return error but returns %v", err)
		}

		if c.exp != act {
			t.Fatalf("expected is %d but actual is %d", c.exp, act)
		}
	}
}

func TestFindFromBehind2(t *testing.T) {
	singleList := list.NewSingleList()
	singleList.Add(1)
	singleList.Add(33)
	singleList.Add(21)
	singleList.Add(33)

	k := 5
	act, err := findFromBehind(singleList, k)

	if err == nil {
		t.Fatalf("should be return the error but is nil and returns %v", act)
	}
}

func TestLength(t *testing.T) {
	singleList := list.NewSingleList()
	singleList.Add(1)
	singleList.Add(33)
	singleList.Add(21)
	singleList.Add(33)

	exp := 4
	act := Length(singleList)

	if exp != act {
		t.Fatalf("expected is %d but actual is %d", exp, act)
	}
}
