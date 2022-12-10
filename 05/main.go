package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			// check whether the  values are not space or [ or ]
			var pos = 1 + 4*s
			var char = lines[i][pos]
			if char != 32 && char != '[' && char != ']' {
				parsed[s] = append(parsed[s], string(char))
			}
		}
	}
	return n, parsed
}

type Ins struct {
	n    int
	from int
	to   int
}

func parseIns(lines []string) []Ins {
	var ins []Ins
	// split the string by spaces
	for _, k := range lines {
		var sp = strings.Split(k, " ")
		if len(sp) != 6 {
			panic("invalid instruction string got")
		}
		n, _ := strconv.Atoi(sp[1])
		f, _ := strconv.Atoi(sp[3])
		t, _ := strconv.Atoi(sp[5])
		i := Ins{
			from: f,
			to:   t,
			n:    n,
		}
		ins = append(ins, i)
	}
	return ins
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
			continue
		}
		if !isIns {
			stack = append(stack, line)
		} else {
			ins = append(ins, line)
		}
	}
	return stack, ins
}

func partA(stack [][]string, instructions []Ins, n int) string {
	for _, ins := range instructions {
		i, j := ins.from-1, ins.to-1
		ssize := len(stack[i])
		// s[start:end] [start, end)
		// elements from start index to end-1
		popped := stack[i][ssize-ins.n : ssize] // get the last n items
		// update the source stack
		stack[i] = stack[i][:ssize-ins.n]
		for k := len(popped) - 1; k >= 0; k-- {
			// as for a stack lifo -> last in first out
			// so insert the last one first and so...on
			stack[j] = append(stack[j], popped[k])
		}
	}
	var res string = ""
	for _, s := range stack {
		res += s[len(s)-1]
	}
	return res
}

func main() {
	var stack, ins = get()
	var n, stacks = parseStack(stack)
	var mv = parseIns(ins)
	var res = partA(stacks, mv, n)
	fmt.Println(res)
}
