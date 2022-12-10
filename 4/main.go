package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a int64
	b int64
}

func parseLine(pair string) [2]Pair {
	var pairs []string = strings.Split(pair, ",")
	var p [2]Pair
	for i := 0; i < len(pairs); i++ {
		var s = strings.Split(pairs[i], "-")
		a, e := strconv.ParseInt(s[0], 10, 64)
		if e != nil {
			fmt.Println("something went wrong ", e)
			os.Exit(1)
		}
		b, e := strconv.ParseInt(s[1], 10, 64)
		if e != nil {
			fmt.Println("something went wrong ", e)
			os.Exit(1)
		}
		p[i].a = a
		p[i].b = b
	}
	return p
}

//  0    1
// a-b, a-b
// 2-4, 4-6
// 6-6, 4-6
// 2-8, 10, 12

func isFullRange(x, y Pair) bool {
	if x.a <= y.a && x.b >= y.b {
		return true
	}
	if x.a >= y.a && x.b <= y.b {
		return true
	}
	return false
}

func isOverlapping(x, y Pair) bool {
	return x.a <= y.b && x.b >= y.a
}

func part2(pairs []string) int {
	var count int = 0
	for _, pair := range pairs {
		var parsed [2]Pair = parseLine(pair)
		if isOverlapping(parsed[0], parsed[1]) {
			fmt.Println("pair", pair)
			count++
		}
	}
	return count
}

func part1(pairs []string) int {
	var count int = 0
	for _, pair := range pairs {
		var parsed [2]Pair = parseLine(pair)
		if isFullRange(parsed[0], parsed[1]) {
			count++
		}
	}
	return count
}

func read() []string {
	fx, e := os.Open("input.data")

	if e != nil {
		println("something went wrong at opening file")
		os.Exit(1)
	}
	reader := bufio.NewReader(fx)
	var pairs []string

	for {
		line, e := reader.ReadString('\n')
		if e != nil {
			break
		} else {
			// remove the newline character
			pairs = append(pairs, line[:len(line)-1])
		}
	}
	return pairs
}

func main() {
	var pairs = read()
	fmt.Println(part1(pairs))
	fmt.Println(part2(pairs))
}
