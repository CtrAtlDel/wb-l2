package main

import (
	"testing"
)

func TestParseFieldIndex(t *testing.T) {
	cases := []struct {
		arg  string
		want int
	}{
		{"", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	for _, c := range cases {
		got := parseFieldIndex(c.arg)
		if got != c.want {
			t.Errorf("parseFieldIndex(%q) == %d, want %d", c.arg, got, c.want)
		}
	}
}
