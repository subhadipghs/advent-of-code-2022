package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMarkerPos(signal string, ws int) int {
	// implementation
	var chars map[string]int = make(map[string]int)
	// map of the first window
	// bvwbjplb
	// first window-> b:2,v:1,w:1
	// first element of the window should decrement by 1
	// now map -> b:1,b:1,w:1
	// incoming char -> j
	// now->  b:1,v:1,w:1,j:1 now check the length of the window which is 4
	// if the length is 4 then return last characters position which is j's position -> 4
	// return the window
	for i := 0; i < ws; i++ {
		var key = string(signal[i])
		chars[key] += 1
		if len(chars) == ws {
			return i + 1
		}
	}
	// now iterate through the rest of string
	// if window size is 4, then iterate through index 4 to rest of the string
	for i := ws; i < len(signal); i++ {
		prev := string(signal[i-ws])
		curr := string(signal[i])
		chars[prev] -= 1 // remove the first element of the previous element
		// if the previous element becomes 0 means
		// it does not exist in the current character window
		// remove that also
		if chars[prev] == 0 {
			delete(chars, prev)
		}
		// increment the new char count
		chars[curr] += 1
		if len(chars) == ws {
			return i + 1
		}
	}
	return -1
}

func solve(lines []string, unq int) int {
	var res = 0
	for _, k := range lines {
		res += getMarkerPos(k, unq)
	}
	return res
}

func main() {
	fx, e := os.Open("input.data")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fx)

	var lines []string
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	var res1 = solve(lines, 4)
	var res2 = solve(lines, 14)
	fmt.Printf("part a - %d, part b - %d\n", res1, res2)
}
