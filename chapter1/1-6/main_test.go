package main

import "testing"

func TestSimpleCompression(t *testing.T) {
	cases := []struct {
		s   string
		exp string
	}{
		{s: "aabcccccaaa", exp: "a2b1c5a3"},
		{s: "abcdefgg", exp: "abcdefgg"},
		{s: "aaaaaaaaaaaa", exp: "a12"},
	}

	for _, c := range cases {
		act := simpleCompression(c.s)
		if act != c.exp {
			t.Fatalf("expected is %v but actual is %v", c.exp, act)
		}
	}
}
