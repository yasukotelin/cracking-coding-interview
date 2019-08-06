package main

import "testing"

func TestReplaceWhiteSpaces(t *testing.T) {
	testCases := []struct {
		s   string
		exp string
	}{
		{s: "hello", exp: "hello"},
		{s: "he llo", exp: "he%20llo"},
		{s: "Mr John Smith", exp: "Mr%20John%20Smith"},
	}

	for _, c := range testCases {
		act := whiteSpacesToPer20(c.s)
		if act != c.exp {
			t.Fatalf("expected is %v but actual is %v", c.exp, act)
		}
	}
}
