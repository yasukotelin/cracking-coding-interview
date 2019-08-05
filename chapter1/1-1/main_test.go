package main

import "testing"

type testCase struct {
	s   string
	exp bool
}

var testCases = []testCase{
	testCase{s: "hello, world", exp: true},
	testCase{s: "fruit", exp: false},
}

func TestHasDuplication1(t *testing.T) {
	for _, c := range testCases {
		actual := hasDuplication1(c.s)
		if actual != c.exp {
			t.Fatalf("expected is %v but actual is %v", c.exp, actual)
		}
	}
}

func TestHasDuplication2(t *testing.T) {
	for _, c := range testCases {
		actual := hasDuplication2(c.s)
		if actual != c.exp {
			t.Fatalf("expected is %v but actual is %v", c.exp, actual)
		}
	}
}

func TestHasDuplication3(t *testing.T) {
	for _, c := range testCases {
		actual := hasDuplication3(c.s)
		if actual != c.exp {
			t.Fatalf("expected is %v but actual is %v", c.exp, actual)
		}
	}
}
