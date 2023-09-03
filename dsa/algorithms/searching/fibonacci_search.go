package main

import (
	"fmt"
	"math"
)

func fibonacciSearch(arr []int, target int) int {
	n := len(arr)
	fibNM2 := 0
	fibNM1 := 1
	fibN := fibNM1 + fibNM2

	for fibN < n {
		fibNM2 = fibNM1
		fibNM1 = fibN
		fibN = fibNM1 + fibNM2
	}

	offset := -1
	for fibN > 1 {
		i := int(math.Min(float64(offset+fibNM2), float64(n-1)))
		if arr[i] < target {
			fibN = fibNM1
			fibNM1 = fibNM2
			fibNM2 = fibN - fibNM1
			offset = i
		} else if arr[i] > target {
			fibN = fibNM2
			fibNM1 = fibNM1 - fibNM2
			fibNM2 = fibN - fibNM1
		} else {
			return i
		}
	}
	if fibNM1 == 0 && arr[offset+1] == target {
		return offset + 1
	}
	return -1
}

func main() {

	arr := []int{10, 20, 30, 60, 80, 160, 800, 980, 1900}
	fmt.Println(fibonacciSearch(arr, 60))

}
