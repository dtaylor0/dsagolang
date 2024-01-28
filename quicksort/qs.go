package main

import (
	"fmt"
)

func partition(lo int, hi int, arr []int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}
	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot
	return idx
}

func quicksort(lo int, hi int, arr []int) {
	if lo >= hi {
		return
	}
	idx := partition(lo, hi, arr)
	quicksort(lo, idx - 1, arr)
	quicksort(idx + 1, hi, arr)
}

func main() {
	test := [][]int{{3, 23, 4, 5, 21, 8}, {9, 8, 7, 6, 3, 2, 4, 1}}
	for idx, arr := range test {
		fmt.Println(idx)
		fmt.Println(arr)
		quicksort(0, len(arr)-1, arr)
		fmt.Println(arr)
	}
}
