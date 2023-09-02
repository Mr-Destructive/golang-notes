package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			swap(&arr[i], &arr[min])
		}
	}
	return arr
}

func main() {

	arr := []int{30, 20, 15, 10}

	fmt.Println(arr)
	SelectionSort(arr)
	fmt.Println(arr)

}
