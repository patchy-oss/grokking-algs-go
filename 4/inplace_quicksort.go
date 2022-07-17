package main

import "fmt"

func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := len(arr) / 2
	cur := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot++
		}
	}

	qsort(arr[:pivot])
	qsort(arr[pivot+1:])
}

func main() {
	arr := []int{1, 0, 3, -1, 2, 11, 10, -2, 0, -1, 10, -5, -10, 10, 10}
	qsort(arr)
	fmt.Println(arr)
}
