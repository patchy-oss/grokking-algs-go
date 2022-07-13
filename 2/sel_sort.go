package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	tmp := []int{1, 8, -1, -13, 0, 2, 0, 1, 3}
	selectionSort(tmp)
	fmt.Println(tmp)
}
