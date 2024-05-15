package main

import "fmt"

func bubble() {

	fmt.Println(bubbleSort([]int{-2, 45, 0, 11, -9}))
	fmt.Println(bubbleSort([]int{4, 5, 6, 7, 8, 9, 11}))
	fmt.Println(bubbleSort([]int{-2, 45, 0, 11, -9, -5, 22, 19, 57, 91, 3, 67}))

}

func bubbleSort(numbers []int) []int {

	preSorted := true //optimizations if array is already sorted

	for k := 0; k < len(numbers); k++ {

		for i := 0; i < (len(numbers) - k); i++ {
			//bounds check
			if (i + 1) < len(numbers) {
				if numbers[i] > numbers[i+1] {
					temp := numbers[i+1]
					numbers[i+1] = numbers[i]
					numbers[i] = temp

					preSorted = false
				}
			}
		}

		if preSorted {
			break
		}
	}
	return numbers
}
