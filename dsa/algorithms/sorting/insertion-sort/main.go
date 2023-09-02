package main

import "fmt"

func InsertionSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n; i++ {
		k := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > k {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = k
	}
	return arr
}

func main() {

	arr := []int{10, 40, 20, 30}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)

}
