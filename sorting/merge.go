package main

import "fmt"

func merge() {

	fmt.Println(mergeSort([]int{6, 5, 12, 10, 9, 2, 1}))
	//fmt.Println(mergeSort([]int{-2, 45, 0, 11, -9}))
	//fmt.Println(mergeSort([]int{4, 5, 6, 7, 8, 9, 11}))
	//fmt.Println(mergeSort([]int{-2, 45, 0, 11, -9, -5, 22, 19, 57, 91, 3, 67}))
}

func mergeSort(numbers []int) []int {

	if len(numbers) > 2 {
		splitResult := ms(numbers[:len(numbers)/2], numbers[len(numbers)/2:])
		fmt.Println(splitResult)
	}

	return numbers
}

func ms(a1 []int, a2 []int) []int {

	var result []int
	//base cases
	if len(a1) == 1 {
		if len(a2) == 1 {
			if a1[0] <= a2[0] {
				return []int{a1[0], a2[0]}
			} else {
				return []int{a2[0], a1[0]}
			}
		}
		if len(a2) > 1 {
			r := []int{}
			r = append(a1, ms(a2[:len(a2)/2], a2[len(a2)/2:])...)
			return r
		}
	}

	if len(a2) == 1 {
		if len(a1) > 1 {
			r := []int{}
			r = append(r, ms(a1[:len(a1)/2], a1[len(a1)/2:])...)
			return r
		}
	}

	var r []int
	if len(a1) > 1 {
		r = append(a2, ms(a1[:len(a1)/2], a1[len(a1)/2:])...)
		return r
	}

	if len(a2) > 1 {
		r = append(a1, ms(a2[:len(a2)/2], a2[len(a2)/2:])...)
		return r
	}

	return result
}
