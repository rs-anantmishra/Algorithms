package main

import "fmt"

func CallCanSum() {

	result := canSum(7, []int{5, 3, 4})
	fmt.Println(result)
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
			for k := range numbers {
				if len(table) > (i + k) {
					table[i+k] = true //v + val
				}
			}
		}
	}

	fmt.Println(table)
	result = table[targetSum-1]
	return result
}
