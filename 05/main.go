package main

import (
	"bufio"
	"fmt"
	"os"
)

const SPACE_CHAR_CODE = 32

func parseStack(lines []string) (int, [][]string) {
	var parsed [][]string
	var n int
	// last line which contains the stack numbering
	last := lines[len(lines)-1]
	for i := 0; i < len(last); i++ {
		if last[i] != SPACE_CHAR_CODE {
			n++
		}
	}
	parsed = make([][]string, n)
	// now slice the last element and iterate through the others
	lines = lines[:len(lines)-1]
	for i := len(lines) - 1; i >= 0; i-- {
		// for each line iterate through the string lines
		// and put the item in specific stack element
		// 1 - 1 // 1 + 4*0
		// 2 - 5 // 1 + 4*1
		// 3 - 9 // 1 + 4*2
		// 4 - 13  // 1 + 4*3
		for s := 0; s < n; s++ {
			// if got space then increment the value
			// otherwise if we found an element
			var pos = 1 + 4*s
			var char = lines[i][pos]
			if char != 32 && char != '[' && char != ']' {
				parsed[s] = append(parsed[s], string(char))
			}
		}
	}
	return n, parsed
}

func get() ([]string, []string) {
	fx, e := os.Open("input.data")
	var stack, ins []string
	var isIns bool = false
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fx)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isIns = true
		}
		if !isIns {
			stack = append(stack, line)
		} else {
			ins = append(ins, line)
		}
	}
	return stack, ins
}

func main() {
	stack, ins := get()
	fmt.Println(stack, ins)
}
