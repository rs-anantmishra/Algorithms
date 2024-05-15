package main

import "fmt"

func selection() {

	fmt.Println(selectionSort([]int{20, 12, 10, 15, 2}))
	fmt.Println(selectionSort([]int{-2, 45, 0, 11, -9}))
	fmt.Println(selectionSort([]int{4, 5, 6, 7, 8, 9, 11}))
	fmt.Println(selectionSort([]int{-2, 45, 0, 11, -9, -5, 22, 19, 57, 91, 3, 67}))
}

func selectionSort(numbers []int) []int {

	var min int
	var minIndex int

	for i := 0; i < len(numbers); i++ {
		min = numbers[i]
		minIndex = i
		for j := i; j < len(numbers); j++ {
			if numbers[j] < min {
				min = numbers[j]
				minIndex = j
			}
		}
		temp := numbers[i]
		numbers[i] = min
		numbers[minIndex] = temp
	}

	return numbers
}
