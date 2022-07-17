package main

import "fmt"

func recBsearch(arr []int, item int) int {
	var iter func(low, high int) int

	iter = func(low, high int) int {
		if low > high {
			return -1
		}

		mid := (low + high) / 2
		if arr[mid] == item {
			return mid
		} else if arr[mid] > item {
			return iter(low, mid-1)
		} else {
			return iter(mid+1, high)
		}
	}

	return iter(0, len(arr)-1)
}

func main() {
	arr100 := make([]int, 100)
	for i := 0; i < len(arr100); i++ {
		arr100[i] = i
	}
	arr1000000000 := make([]int, 1000000000)
	randArr := []int{-2, -1, 0, 1, 5, 8, 10, 30, -3}

	for i := 0; i < len(arr1000000000); i++ {
		arr1000000000[i] = i
	}

	res := recBsearch(arr100, 49)
	fmt.Println("100: ", res)
	res = recBsearch(arr1000000000, 7912627)
	fmt.Println("1000000000: ", res)
	res = recBsearch(randArr, -3)
	fmt.Println("rand: ", res)
	res = recBsearch(arr100, 101)
	fmt.Println("incorrect: ", res)
}
