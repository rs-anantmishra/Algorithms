package main

import "fmt"

func CallGridTraveler() {

	result := gridTraveler(10, 2)
	fmt.Println("answer is:", result)
}

func gridTraveler(cols int, rows int) int {

	//add an extra position of 0 so that indices align with inputs
	//account for OFF BY 1 NATURE
	cols++
	rows++

	//create a table - in this case an x*y size table as a 2d array/slice
	//initialize all values to zero
	var table [][]int

	for i := rows; i > 0; i-- {
		temp := make([]int, cols)
		for i := range rows {
			temp[i] = 0
		}
		table = append(table, temp)
	}

	//base case - seed value - if grid is 1x1 size there is only 1 way to traverse it.
	table[1][1] = 1

	//starting at 0,0 add the current value to right and down neighbours (with bourds checks)
	for i := 0; i < rows; i++ {
		for k := 0; k < cols; k++ {

			//increment next rows-cells with bounds check
			if i+1 < rows {
				table[i+1][k] = table[i+1][k] + table[i][k]
			}

			//increment next cols-cells with bounds check
			if k+1 < cols {
				table[i][k+1] = table[i][k+1] + table[i][k]
			}
		}
	}

	return table[rows-1][cols-1]
}
