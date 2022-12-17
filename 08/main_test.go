package main

import "testing"

func Test(t *testing.T) {

	t.Run("should return correct number of visible trees", func(t *testing.T) {
		want := 21
		inp := [][]int64{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		}
		got := solve(inp)
		if want != got {
			t.Errorf("opps! failed to return correct number of visible trees! got - %d, want - %d", got, want)
		}
	})

}
