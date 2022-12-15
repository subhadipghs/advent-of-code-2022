package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	label     string
	isDir     bool
	size      int64
	parent    *node
	childrens []*node
}

type tree struct {
	root *node
	size int64
}

func bfs(key string, root *node) *node {
	if root == nil {
		panic("empty root node")
	}
	var queue []*node
	var disc map[*node]bool = make(map[*node]bool)
	queue = append(queue, root)

	for len(queue) > 0 {
		front := queue[0] // get the first element
		queue = queue[1:] // pop the first element
		disc[front] = true
		// for each of the childrens
		for _, k := range front.childrens {
			// if the label found then just do it
			if k.label == key {
				return k
			}
			if !disc[k] {
				disc[k] = true
				queue = append(queue, k)
			}
		}
	}
	return nil
}

// something I would like to introduce is naive tree
func makeNode(label string, size int64, isDir bool, parent *node) *node {
	return &node{
		label:     label,
		childrens: make([]*node, 0),
		size:      size,
		isDir:     isDir,
		parent:    parent,
	}
}

func (n *node) add(ch *node) int64 {
	if n == nil {
		fmt.Println("oops! nil node")
		return 0
	}
	if !n.isDir {
		fmt.Println(n)
		panic("not a directory! cannot add a child node to this")
	}
	if ch == nil {
		panic("oops! invalid child node")
	}
	n.childrens = append(n.childrens, ch)
	var q = n
	for q != nil {
		q.size += ch.size
		q = q.parent
	}
	return n.size
}

func split(line string) []string {
	return strings.Split(line, " ")
}

// PART 1
func getDirsPart1(root *node) int64 {
	var totalSize int64 = 0

	if root == nil {
		panic("empty root node")
	}
	var queue []*node
	var disc map[*node]bool = make(map[*node]bool)
	queue = append(queue, root)

	for len(queue) > 0 {
		front := queue[0] // get the first element
		queue = queue[1:] // pop the first element
		disc[front] = true
		// for each of the childrens
		for _, k := range front.childrens {
			// if the label found then just do it
			if !disc[k] {
				if k.isDir && k.size <= 100000 {
					totalSize += k.size
				}
				disc[k] = true
				queue = append(queue, k)
			}
		}
	}
	return totalSize
}

func parseCmds(lines []string) *tree {
	var fs *tree = &tree{
		root: makeNode("/", 0, true, nil),
		size: 0,
	}
	var curr *node

	for _, k := range lines {
		var s []string = split(k)
		// check whether it's a command or not
		if s[0] == "$" {
			cmd := s[1]
			if cmd == "cd" {
				var dir = s[2]
				switch dir {
				case "/":
					curr = fs.root
				case "..":
					curr = curr.parent
				default:
					curr = bfs(dir, curr)
				}
			}
		} else {
			// here is the result of the ls command
			// if starts with dir then it's a directory otherwise it's a file
			var name = s[1]
			if s[0] == "dir" {
				dir := makeNode(name, 0, true, curr)
				curr.add(dir)
			} else {
				var size, _ = strconv.ParseInt(s[0], 10, 64)
				var file = makeNode(name, size, false, curr)
				curr.add(file)
			}
		}
	}
	return fs
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
		lines = append(lines, scanner.Text())
	}
	fs := parseCmds(lines)
	fmt.Println(getDirsPart1(fs.root))
}
