package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func split(line string) []int64 {
	var s = strings.Split(line, "")
	var n []int64
	for _, k := range s {
		var i, _ = strconv.ParseInt(k, 10, 64)
		n = append(n, i)
	}
	return n
}

func parse(lines []string) [][]int64 {
	var arr [][]int64
	for _, l := range lines {
		var row = split(l)
		arr = append(arr, row)
	}
	return arr
}

func solve2(inp [][]int64) int {
	var score int = 0
	// 1. count if the it's a boundary element
	// 2. count if the all the other row and column elements are lesser than that
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[i]); j++ {
			// if it's the first row or last row dont' count
			// if it's the first column or last column dont count
			if i != 0 && i != len(inp)-1 && j != 0 && j != len(inp)-1 {
				// if every tree between the edge and the tree itself is smaller
				// or conversly if any tree between the edge and tree is same or bigger height than it then it's not visible
				// a tree can be seen from 4 sides, right, left, up, down
				scores := []int{
					0,
					0,
					0,
					0,
				}
				// from left
				for k := j - 1; k >= 0; k-- {
					scores[0]++
					if inp[i][k] >= inp[i][j] {
						break
					}
				}
				// from right
				for k := j + 1; k < len(inp[i]); k++ {
					scores[1]++
					if inp[i][k] >= inp[i][j] {
						break
					}
				}
				// from top
				for k := i - 1; k >= 0; k-- {
					scores[2]++
					if inp[k][j] >= inp[i][j] {
						break
					}
				}
				// from down
				for k := i + 1; k < len(inp); k++ {
					scores[3]++
					if inp[k][j] >= inp[i][j] {
						break
					}
				}
				// multiply them together
				local := scores[0] * scores[1] * scores[2] * scores[3]
				if score < local {
					score = local
				}
			}
		}
	}
	return score
}

func solve(inp [][]int64) int {
	var res int = 0
	// 1. count if the it's a boundary element
	// 2. count if the all the other row and column elements are lesser than that
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[i]); j++ {
			// if it's the first row or last row dont' count
			// if it's the first column or last column dont count
			if (i == 0 || i == len(inp)-1) || (j == 0 || j == len(inp)-1) {
				res++
			} else {
				// if every tree between the edge and the tree itself is smaller
				// or conversly if any tree between the edge and tree is same or bigger height than it then it's not visible
				// a tree can be seen from 4 sides, right, left, up, down
				visible := []bool{
					true,
					true,
					true,
					true,
				}
				// from left
				for k := 0; k < j; k++ {
					if inp[i][k] >= inp[i][j] {
						visible[0] = false
						break
					}
				}
				// from right
				for k := j + 1; k < len(inp[i]); k++ {
					if inp[i][k] >= inp[i][j] {
						visible[1] = false
						break
					}
				}
				// from top
				for k := 0; k < i; k++ {
					if inp[k][j] >= inp[i][j] {
						visible[2] = false
						break
					}
				}

				// from down
				for k := i + 1; k < len(inp); k++ {
					if inp[k][j] >= inp[i][j] {
						visible[3] = false
						break
					}
				}
				// if any one side is true then it's visible
				for _, k := range visible {
					if k {
						res++
						break
					}
				}
			}
		}
	}
	return res
}

func main() {
	var fx, e = os.Open("input.actual")

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	var scanner = bufio.NewScanner(fx)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	inp := parse(lines)
	// part 1
	fmt.Println(solve(inp))
	// part 2
	fmt.Println(solve2(inp))
}
