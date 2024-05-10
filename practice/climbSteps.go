/*
You are Climing the staircase. it takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top.
*/

package main

import "fmt"

func Steps() {

	totalSteps := 50
	memo := map[int]int{}

	fmt.Println(calculateSteps(totalSteps, memo))

	fmt.Println("neetSteps", neetSteps(totalSteps))

}

func calculateSteps(totalSteps int, memo map[int]int) int {

	//base cases
	if totalSteps == 0 {
		return 0
	}
	if totalSteps == 1 {
		return 1
	}
	if totalSteps == 2 {
		return 2
	}
	//fast-return
	elem, ok := memo[totalSteps]
	if ok {
		return elem
	}

	//update memo
	memo[totalSteps] = calculateSteps(totalSteps-1, memo) + calculateSteps(totalSteps-2, memo)

	return memo[totalSteps]
}

func neetSteps(totalSteps int) int {
	one, two := 1, 1

	for i := 0; i < (totalSteps - 1); i++ {
		temp := one
		one = one + two
		two = temp
	}
	return one
}
