package main

import "fmt"

func CallHowSum(args []string) {

	var result []int
	targetSum := 344
	numbers := []int{7, 14}

	if args[2] == "m" {
		memo := make(map[int][]int)
		result = memoHowSum(targetSum, numbers, memo)
		fmt.Println(result)
	} else {
		result = howSum(targetSum, numbers)
		fmt.Println(result)
	}

}

func howSum(targetSum int, numbers []int) []int {
	result := []int{}

	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, value := range numbers {
		remainder := targetSum - value

		result = howSum(remainder, numbers)
		if result != nil {
			result = append(result, value)
			return result
		}

	}

	return result
}

func memoHowSum(targetSum int, numbers []int, memo map[int][]int) []int {
	result := []int{}

	//base cases
	if targetSum == 0 {
		return result
	}
	if targetSum < 0 {
		return nil
	}

	//fast-return
	elem, ok := memo[targetSum]
	if ok {
		return elem
	}

	for _, value := range numbers {
		remainder := targetSum - value

		result = memoHowSum(remainder, numbers, memo)
		if result != nil {
			//update memo
			memo[remainder] = result
			result = append(result, value)
			return result
		}
	}

	//update memo
	memo[targetSum] = result
	return result
}
