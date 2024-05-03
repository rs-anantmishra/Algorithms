package main

import (
	"fmt"
)

func CallCanSum(args []string) {

	var result bool

	sum := 7
	numbers := []int{5, 3, 4, 7}

	//numbers := []int{7, 14}
	//sum := 300

	//numbers := []int{2, 3}
	//sum := 7

	//numbers := []int{2, 4}
	//sum := 7

	//numbers := []int{2, 3, 5}
	//sum := 8

	//numbers := []int{7, 14}
	//sum := 7

	memo := make(map[int]bool)
	if args[2] == "m" {
		result = MemoCanSum(sum, numbers, memo)
	} else {
		result = canSum(sum, numbers)
	}

	fmt.Println("Can sum?", result)
}

func MemoCanSum(targetSum int, numbers []int, memo map[int]bool) bool {

	//base cases
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	//fast-return from memo
	answer, ok := memo[targetSum]
	if ok {
		return answer
	}

	var remainder int
	for _, value := range numbers {
		remainder = targetSum - value
		if MemoCanSum(remainder, numbers, memo) {
			memo[targetSum] = true
			return memo[targetSum]
		}
	}

	memo[targetSum] = false
	return memo[targetSum]
}

func canSum(sum int, numbers []int) bool {

	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}

	for i := 0; i < len(numbers); i++ {
		result := sum - numbers[i]

		//since recursive call is not a part of the return statement,
		//a mechanism to return from here is needed separately.
		ans := canSum(result, numbers)
		if ans {
			return true
		}
	}

	return false
}
