package main

import "fmt"

func CallCanSum() {

	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(7, []int{2, 4}))
	fmt.Println(canSum(8, []int{2, 3, 5}))
	fmt.Println(canSum(300, []int{7, 14}))
}

func canSum(targetSum int, numbers []int) bool {
	result := false

	//account for OFF BY 1 NATURE to match indices and values.
	targetSum++

	//create table - 1D table/array
	table := make([]bool, targetSum)

	//init table with zero-fill
	for i := range table {
		table[i] = false
	}

	//seed values - when pointer is on 0th, we have seed values in numbers array
	//we are going to add values to get multiples
	for _, value := range numbers {
		//bounds check
		if len(table) >= value {
			table[value] = true //value
		}
	}

	//if seed value is available at current index
	//if its in bounds - the value at index[current index + val] will be:
	//value at current index + how far we are looking ahead by
	for i, v := range table {
		if v {
			for _, val := range numbers {
				if len(table) > (i + val) {
					table[i+val] = true //v + val
				}
			}
		}
	}

	result = table[targetSum-1]
	return result
}
