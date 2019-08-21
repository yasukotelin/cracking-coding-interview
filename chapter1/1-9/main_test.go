package main

import "testing"

func TestIsCircledString(t *testing.T) {
	cases := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{s1: "waterbottle", s2: "erbottlewat", exp: true},
		{s1: "yasukotelin", s2: "linyasukote", exp: true},
		{s1: "hello, world", s2: "world, hello", exp: false},
	}

	for _, c := range cases {
		if isCircledString(c.s1, c.s2) != c.exp {
			t.Fatalf("Fatal: s1: %s s2: %s", c.s1, c.s2)
		}
	}

}
