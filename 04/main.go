package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	start int64
	end   int64
}

func parseLine(pair string) [2]Pair {
	var pairs []string = strings.Split(pair, ",")
	var p [2]Pair
	for i := 0; i < len(pairs); i++ {
		var s = strings.Split(pairs[i], "-")
		start, e := strconv.ParseInt(s[0], 10, 64)
		if e != nil {
			fmt.Println("something went wrong ", e)
			os.Exit(1)
		}
		end, e := strconv.ParseInt(s[1], 10, 64)
		if e != nil {
			fmt.Println("something went wrong ", e)
			os.Exit(1)
		}
		p[i].start = start
		p[i].end = end
	}
	return p
}

//  0    1
// start-end, start-end
// 2-4, 4-6
// 6-6, 4-6
// 2-8, 10, 12

func isFullRange(x, y Pair) bool {
	if x.start <= y.start && x.end >= y.end {
		return true
	}
	if x.start >= y.start && x.end <= y.end {
		return true
	}
	return false
}

func isOverlapping(x, y Pair) bool {
	if x.start > y.start {
		x.start, x.end, y.start, y.end = y.start, y.end, x.start, x.end
	}
	return y.start <= x.end && y.start >= x.start
}

func part2(pairs []string) int {
	var count int = 0
	for _, pair := range pairs {
		var parsed [2]Pair = parseLine(pair)
		if isOverlapping(parsed[0], parsed[1]) {
			//fmt.Println(pair)
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
	reader := bufio.NewScanner(fx)
	var pairs []string

	for reader.Scan() {
		line := reader.Text()
		pairs = append(pairs, line)
	}
	return pairs
}

func main() {
	var pairs = read()
	fmt.Println(len(pairs))
	fmt.Println(part2(pairs))
}
