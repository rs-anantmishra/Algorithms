package main

import "fmt"

func CallHowSum(args []string) {

	targetSum := 8
	numbers := []int{5, 3, 4, 3}

	result := howSum(targetSum, numbers)
	fmt.Println(result)

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
