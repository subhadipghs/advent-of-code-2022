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
}
