package main

import "testing"

func Test(t *testing.T) {
	t.Run("should return all completing pairs", func(t *testing.T) {
		got := part1([]string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"2-8,3-7",
			"6-6,4-6",
			"2-6,4-8",
			"2-8,10-12",
		})

		want := 2

		if got != want {
			t.Errorf("expected - %d, returned - %d", want, got)
		}
	})

	t.Run("should parse the string correctly", func(t *testing.T) {
		got := parseLine("2-4,6-8")
		want := [2]Pair{
			{a: 2, b: 4},
			{a: 6, b: 8},
		}
		for i := 0; i < len(want); i++ {
			if got[i].a != want[i].a || got[i].b != want[i].b {
				t.Errorf("got - %d-%d, want - %d-%d", got[i].a, got[i].b, want[i].a, want[i].b)
			}
		}
	})

	t.Run("should return correct number of overlapping ranges", func(t *testing.T) {
		var pairs = read()
		got := part2(pairs)

		want := 4

		if got != want {
			t.Errorf("expected - %d, returned - %d", want, got)
		}
	})
}
