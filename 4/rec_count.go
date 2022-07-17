package main

import "fmt"

type node struct {
	val  int
	next *node
}

func recCount(head *node) int {
	if head == nil {
		return 0
	}

	return 1 + recCount(head.next)
}

func main() {
	list := &node{1, &node{2, &node{3, nil}}}
	fmt.Println(recCount(list))
	var empty *node = nil
	fmt.Println(recCount(empty))
}
