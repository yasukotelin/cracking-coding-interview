package main

import "testing"

func TestIsSame(t *testing.T) {
	testCases := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{
			s1:  "hello",
			s2:  "world",
			exp: false,
		},
		{
			s1:  "hello",
			s2:  "elolh",
			exp: true,
		},
		{
			s1:  "Dog",
			s2:  "dog",
			exp: false,
		},
		{
			s1:  "Golang ",
			s2:  "gnalog",
			exp: false,
		},
	}

	for _, c := range testCases {
		act := isSame(c.s1, c.s2)
		if act != c.exp {
			t.Fatalf("s1: %s s2: %s, expected is %v but actual %v", c.s1, c.s2, c.exp, act)
		}
	}
}
