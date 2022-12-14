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
	size      int
	parent    *node
	childrens []*node
}

type tree struct {
	root *node
	size int
}

func bfs(key string, root *node) *node {
	if root == nil {
		panic("empty root node")
	}
	var queue []*node
	var disc map[string]bool = make(map[string]bool)
	queue = append(queue, root)

	for len(queue) > 0 {
		front := queue[0] // get the first element
		queue = queue[1:] // pop the first element
		disc[front.label] = true
		// for each of the childrens
		for _, k := range front.childrens {
			// if the label found then just do it
			if k.label == key {
				return k
			}
			if !disc[k.label] {
				disc[k.label] = true
				queue = append(queue, k)
			}
		}
	}
	return nil
}

// something I would like to introduce is naive tree
func makeNode(label string, size int, isDir bool, parent *node) *node {
	return &node{
		label:     label,
		childrens: make([]*node, 0),
		size:      size,
		isDir:     isDir,
		parent:    parent,
	}
}

func (n *node) add(ch *node) int {
	if n == nil {
		fmt.Println("oops! nil node")
		return 0
	}
	if !n.isDir {
		panic("not a directory! cannot add a child node to this")
	}
	if ch == nil {
		panic("oops! invalid child node")
	}
	var size = n.size + ch.size
	n.childrens = append(n.childrens, ch)
	n.size = size
	return size
}

func printTree(n *node) {
	if n == nil {
		return
	} else {
		var t string
		if n.isDir {
			t = "dir"
		} else {
			t = "file"
		}
		fmt.Printf("%s (%s) - %d\n", n.label, t, n.size)
		if len(n.childrens) > 0 {
			for i := 0; i < len(n.childrens); i++ {
				printTree(n.childrens[i])
			}
		}
	}
}

func split(line string) []string {
	return strings.Split(line, " ")
}

func parseCmds(lines []string) {
	var fs *tree
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
					fs = &tree{
						root: makeNode("/", 0, true, nil),
						size: 0,
					}
					curr = fs.root
				case "..":
					// go back to parent directory
					if curr.parent != nil {
						// fmt.Println("going to parent dir of -", curr.label)
						curr = curr.parent
						// fmt.Println("current directory now -", curr.label)
					} else {
						// why is current's parent directory nil here?
						panic("current parent is nil here")
					}
				default:
					// otherwise move to the respective directory
					// fmt.Println("finding dir -", dir)
					curr = bfs(dir, fs.root)
					// if curr != nil {
					// fmt.Println("result of finding", curr.label)
					// } else {
					// fmt.Println("could not find the directory")
					// }
				}
			} else if cmd == "ls" {
				// list command
			} else {
				panic("unknown command")
			}
		} else {
			// here is the result of the ls command
			// if starts with dir then it's a directory otherwise it's a file
			var name = s[1]
			if s[0] == "dir" {
				dir := makeNode(name, 0, true, curr)
				curr.add(dir)
			} else {
				var size, _ = strconv.ParseInt(s[0], 10, 32)
				var file = makeNode(name, int(size), false, curr)
				curr.add(file)
			}
		}
	}
	printTree(fs.root)
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
	parseCmds(lines)
}
