package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func BubbleSortDesc(arr []int, asc bool) []int {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if asc {
				if arr[j] < arr[i] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			} else {
				if arr[j] > arr[i] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}
		}
	}
	return arr
}

func main() {
	arr := []int{20, 10, 30, 5}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
	BubbleSortDesc(arr, false)
	fmt.Println(arr)

}
