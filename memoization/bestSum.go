package main

import "fmt"

func CallBestSum(args []string) {

	targetSum := 343
	numbers := []int{7, 14}

	//result := bestSum(targetSum, numbers)
	//result := allSum(targetSum, numbers)
	memo := map[int][]int{}
	result := memoBestSum(targetSum, numbers, memo)
	fmt.Println(result)

}

func bestSum(targetSum int, numbers []int) []int {

	var shortestCombination []int
	var combination []int

	//base cases
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, value := range numbers {
		remainder := targetSum - value
		result := bestSum(remainder, numbers)

		if result != nil {
			combination = append(result, value)
			if len(shortestCombination) == 0 || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}
	return shortestCombination
}

func memoBestSum(targetSum int, numbers []int, memo map[int][]int) []int {

	var shortestCombination []int
	var combination []int

	//base cases
	if targetSum == 0 {
		return []int{}
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
		result := memoBestSum(remainder, numbers, memo)

		if result != nil {
			combination = append(result, value)
			if len(shortestCombination) == 0 || len(combination) < len(shortestCombination) {
				memo[targetSum] = combination
				shortestCombination = combination
			}
		}
	}

	memo[targetSum] = shortestCombination
	return memo[targetSum]
}

func allSum(targetSum int, numbers []int) []int {

	var allCombinations [][]int
	var combinations []int

	//base cases
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, value := range numbers {
		remainder := targetSum - value
		result := allSum(remainder, numbers)

		if result != nil {
			combinations = append(result, value)
			allCombinations = append(allCombinations, combinations)
			fmt.Println(allCombinations)

			//if len(allCombinations) == 0 || len(combinations) < len(allCombinations) {
			//	allCombinations = combinations
			//}
		}
	}

	return combinations
}
