package main

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("should parse the stack correctly", func(t *testing.T) {
		var input = []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
		}
		n, stack := parseStack(input)
		nx := 3
		stackx := [][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		}
		if n != nx {
			t.Errorf("[length] got - %d, want - %d", n, nx)
			t.FailNow()
		}

		for i := 0; i < nx; i++ {
			for j := 0; j < len(stackx[i]); j++ {
				if stack[i][j] != stackx[i][j] {
					t.Errorf("unmatched element at index {%d,%d}", i, j)
					t.Error("got", stack)
					t.Error("expected", stackx)
					t.FailNow()
				}
			}
		}
	})

	t.Run("should return the top crates in all stacks", func(t *testing.T) {
		var stacks = [][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		}
		var ins = []Ins{
			{from: 2, to: 1, n: 1},
			{from: 1, to: 3, n: 3},
			{from: 2, to: 1, n: 2},
			{from: 1, to: 2, n: 1},
		}
		var got = partA(stacks, ins, len(stacks))
		var want = "CMZ"
		if got != want {
			t.Errorf("got - %s, want - %s", got, want)
		}
	})

	t.Run("should parse the movement instruction correctly", func(t *testing.T) {
		var input = []string{
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}

		var got = parseIns(input)
		var want = []Ins{
			{from: 2, to: 1, n: 1},
			{from: 1, to: 3, n: 3},
			{from: 2, to: 1, n: 2},
			{from: 1, to: 2, n: 1},
		}

		if len(got) != len(want) {
			t.Error("length does not match")
		}

		for i := 0; i < len(got); i++ {
			var g, w = got[i], want[i]
			if g.from != w.from || g.to != w.to || g.n != w.n {
				t.Errorf("mismatch got and want")
				t.Error("got", g)
				t.Error("want", w)
				t.FailNow()
			}
		}
	})
}
