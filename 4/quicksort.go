package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	lesser := make([]int, 0, len(arr)-1)
	higher := make([]int, 0, len(arr)-1)

	for i := 1; i < len(arr); i++ {
		if arr[i] >= pivot {
			higher = append(higher, arr[i])
		} else {
			lesser = append(lesser, arr[i])
		}
	}

	// turns out it's really hard to read after several months...
	// we sort lesser part, append pivot to it and append higher at the end
	lesser = quicksort(lesser)
	higher = quicksort(higher)
	res := append(append(lesser, pivot), higher...)

	return res
}

func main() {
	arr := []int{1, 0, 3, -1, 2, 11, 10, -2, 0, -1, 10, -5, -10, 10, 10}
	arr = quicksort(arr)
	fmt.Println(arr)
}
