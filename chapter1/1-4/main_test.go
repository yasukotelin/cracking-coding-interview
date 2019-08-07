package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		s   string
		exp bool
	}{
		{s: "Tact Coa", exp: true},
		{s: "Hello", exp: false},
		{s: "mahamh", exp: true},
		{s: "Amahamh", exp: true},
	}

	for _, c := range testCases {
		act := isPalindrome(c.s)
		if act != c.exp {
			t.Fatalf("string is %v: expected is %v but actual is %v", c.s, c.exp, act)
		}
	}
}
