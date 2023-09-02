package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivotIdx := partition(arr, low, high)
	QuickSort(arr, low, pivotIdx-1)
	QuickSort(arr, pivotIdx+1, high)

}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx++
			temp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = temp
		}
	}

	idx++
	arr[high] = arr[idx]
	arr[idx] = pivot

	return idx
}

func main() {
	arr := []int{10, 40, 30, 20}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
