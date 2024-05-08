package main

import (
	"fmt"
	"strings"
)

func CallCountConstruct(args []string) {

	memo := map[string]int{}

	targetWord := "abcdef"
	wordbank := []string{"ab", "abc", "cd", "def", "abcd"}
	fmt.Println("abcdef", countConstruct(targetWord, wordbank))
	fmt.Println("abcdef memoized", memoCountConstruct(targetWord, wordbank, memo))

	targetWord = "stakeboard"
	wordbank = []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}
	fmt.Println("stakeboard", countConstruct(targetWord, wordbank))
	fmt.Println("stakeboard memoized", memoCountConstruct(targetWord, wordbank, memo))

	targetWord = "enterapotentpot"
	wordbank = []string{"a", "p", "ent", "enter", "ot", "o", "t"}
	fmt.Println("enterapotentpot", countConstruct(targetWord, wordbank))
	fmt.Println("enterapotentpot memoized", memoCountConstruct(targetWord, wordbank, memo))

	targetWord = "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef"
	wordbank = []string{"e", "ee", "eee", "eeee", "eeeeeeeeeeeeee"}
	//fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", countConstruct(targetWord, wordbank))
	fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef memoized", memoCountConstruct(targetWord, wordbank, memo))

	targetWord = "purple"
	wordbank = []string{"purp", "p", "ur", "le"}
	fmt.Println("purple", countConstruct(targetWord, wordbank))
	fmt.Println("purple memoized", memoCountConstruct(targetWord, wordbank, memo))

}

func countConstruct(targetWord string, wordbank []string) int {

	result := 0
	//base case
	if targetWord == "" {
		result++
		return result
	}

	for _, v := range wordbank {
		if strings.Index(targetWord, v) == 0 {
			suffix, _ := strings.CutPrefix(targetWord, v)
			if canConstruct(suffix, wordbank) {
				result++
			}
		}
	}

	return result
}

func memoCountConstruct(targetWord string, wordbank []string, memo map[string]int) int {

	result := 0
	//base case
	if targetWord == "" {
		result++
		return result
	}
	//fast-return
	_, ok := memo[targetWord]
	if ok {
		return memo[targetWord]
	}

	for _, v := range wordbank {
		if strings.Index(targetWord, v) == 0 {
			suffix, _ := strings.CutPrefix(targetWord, v)
			if memoCountConstruct(suffix, wordbank, memo) > 0 {
				result++
				memo[targetWord] = result
			}
		}
	}

	memo[targetWord] = result
	return memo[targetWord]
}
