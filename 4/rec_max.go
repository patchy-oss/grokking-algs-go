package main

import "fmt"

type node struct {
	val  int
	next *node
}

func rec_max(head *node) (int, error) {
	var max int
	if head == nil {
		return max, fmt.Errorf("list is empty")
	}

	max, err := rec_max(head.next)
	if err != nil || max <= head.val {
		return head.val, nil
	}

	return max, nil
}

func main() {
	list := &node{4, &node{2, &node{-3, &node{1, nil}}}}
	fmt.Println(rec_max(list))
}
