package main

import "fmt"

func insertion() {

	fmt.Println(insertionSort([]int{9, 5, 1, 4, 3}))
	//fmt.Println(insertionSort([]int{20, 12, 10, 15, 2}))
	//fmt.Println(insertionSort([]int{-2, 45, 0, 11, -9}))
	//fmt.Println(insertionSort([]int{4, 5, 6, 7, 8, 9, 11}))
	//fmt.Println(insertionSort([]int{-2, 45, 0, 11, -9, -5, 22, 19, 57, 91, 3, 67}))
}

func insertionSort(numbers []int) []int {

	for i := 0; i < len(numbers); i++ {

		var nextMin int
		if (i + 1) < len(numbers) {
			nextMin = numbers[i+1]
			minIndex := i + 1
			minReset := false

			for j := minIndex; (j > 0) && !minReset; j-- {
				if nextMin < numbers[j-1] && j-1 >= 0 {
					//swap
					temp := numbers[j-1]
					numbers[j-1] = numbers[j]
					numbers[j] = temp
				} else {
					minReset = !minReset
				}
			}
		}
	}
	return numbers
}
