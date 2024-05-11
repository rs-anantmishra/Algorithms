package main

import "fmt"

func CallHowSum() {

	//fmt.Println(howSum(7, []int{2, 3}))
	//fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	//fmt.Println(howSum(7, []int{2, 4}))
	fmt.Println(howSum(6, []int{2, 3, 5}))
	//fmt.Println(howSum(300, []int{7, 14}))
}

func howSum(targetSum int, numbers []int) []int {

	targetSum++
	table := make([][]int, targetSum)

	//seed 0th value
	table[0] = []int{0}

	for i, v := range table {
		//fmt.Println(table)
		if v != nil {
			for _, val := range numbers {
				if len(table) > (i + val) {

					//seed values at i == 0
					// if (len(table[i+val]) > len(append(v, val))) || i == 0 || table[i+val] == nil {
					if i == 0 || table[i+val] == nil {
						table[i+val] = append(table[i+val], v...)
						table[i+val] = append(table[i+val], val)
					}

				}
			}
		}
	}

	return table[targetSum-1]

}
