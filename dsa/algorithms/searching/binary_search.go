package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{10, 20, 40, 70}
	num := 20
	index := binarySearch(arr, num)
	if index == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found at index", index)
	}
}
