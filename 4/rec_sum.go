package main

import "fmt"

func recSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + recSum(arr[1:])
}

func main() {
	arr := []int{1, -1, 2, 3, -5, 42}
	fmt.Println(recSum(arr))
}
