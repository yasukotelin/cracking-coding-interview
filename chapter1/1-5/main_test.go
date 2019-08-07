package main

import "testing"

func TestCanConvert(t *testing.T) {
	cases := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{s1: "pale", s2: "ple", exp: true},
		{s1: "pales", s2: "pale", exp: true},
		{s1: "pale", s2: "bale", exp: true},
		{s1: "pale", s2: "bake", exp: false},
	}

	for _, c := range cases {
		act := canConvert(c.s1, c.s2)
		if act != c.exp {
			t.Fatalf("expected is %v but actual is %v. s1: %v s2: %v", c.exp, act, c.s1, c.s2)
		}
	}
}
