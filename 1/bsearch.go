package main

import "fmt"

func bsearch(arr []int, item int) (int, int) {
	stp_num := 0
	high := len(arr) - 1
	low := 0
	for low <= high {
		stp_num++
		mid := (high + low) / 2
		guess := arr[mid]
		if guess == item {
			return mid, stp_num
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, stp_num
}

func main() {
	arr100 := make([]int, 100)
	arr1000000000 := make([]int, 1000000000)
	randArr := []int{0, 1, 5, 8, 10, 30}

	for i := 0; i < len(arr100); i++ {
		arr100[i] = i
	}

	for i := 0; i < len(arr1000000000); i++ {
		arr1000000000[i] = i
	}

	res, stp := bsearch(arr100, 1)
	fmt.Println("100: ", res, stp)
	res, stp = bsearch(arr1000000000, 7912627)
	fmt.Println("1000000000: ", res, stp)
	res, stp = bsearch(randArr, 8)
	fmt.Println("rand: ", res, stp)
	res, stp = bsearch(arr100, 101)
	fmt.Println("incorrect: ", res, stp)
}
