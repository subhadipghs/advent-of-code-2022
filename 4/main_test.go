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
		})

		want := 2

		if got != want {
			t.Errorf("expected - %d, returned - %d", want, got)
		}
	})

	t.Run("should parse the string correctly", func(t *testing.T) {
		got := parseLine("2-4,6-8")
		want := [2]Pair{
			{start: 2, end: 4},
			{start: 6, end: 8},
		}
		for i := 0; i < len(want); i++ {
			if got[i].start != want[i].start || got[i].end != want[i].end {
				t.Errorf("got - %d-%d, want - %d-%d", got[i].start, got[i].end, want[i].start, want[i].end)
			}
		}
	})

	t.Run("should return correct number of overlapping ranges", func(t *testing.T) {
		got := part2([]string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"2-8,3-7",
			"6-6,4-6",
			"2-6,4-8",
		})

		want := 4

		if got != want {
			t.Errorf("expected - %d, returned - %d", want, got)
		}
	})
}
