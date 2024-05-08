package main

import (
	"fmt"
	"strings"
)

func CallCanConstruct(args []string) {

	memo := map[string]bool{}

	targetWord := "abcdef"
	wordbank := []string{"ab", "abc", "cd", "def", "abcd"}
	fmt.Println("abcdef", canConstruct(targetWord, wordbank))
	fmt.Println("abcdef memoized", memoCanConstruct(targetWord, wordbank, memo))

	targetWord = "stakeboard"
	wordbank = []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}
	fmt.Println("stakeboard", canConstruct(targetWord, wordbank))
	fmt.Println("stakeboard memoized", memoCanConstruct(targetWord, wordbank, memo))

	targetWord = "enterapotentpot"
	wordbank = []string{"a", "p", "ent", "enter", "ot", "o", "t"}
	fmt.Println("enterapotentpot", canConstruct(targetWord, wordbank))
	fmt.Println("enterapotentpot memoized", memoCanConstruct(targetWord, wordbank, memo))

	targetWord = "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef"
	wordbank = []string{"e", "ee", "eee", "eeee", "eeeeeeeeeeeeee"}
	//fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", canConstruct(targetWord, wordbank))
	fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef memoized", memoCanConstruct(targetWord, wordbank, memo))

}

func canConstruct(targetWord string, wordbank []string) bool {

	//base case
	if targetWord == "" {
		return true
	}

	for _, v := range wordbank {
		if strings.Index(targetWord, v) == 0 {
			suffix, _ := strings.CutPrefix(targetWord, v)
			if canConstruct(suffix, wordbank) {
				return true
			}
		}
	}

	return false
}

func memoCanConstruct(targetWord string, wordbank []string, memo map[string]bool) bool {

	//base case
	if targetWord == "" {
		return true
	}
	//fast-return
	_, ok := memo[targetWord]
	if ok {
		return memo[targetWord]
	}

	for _, v := range wordbank {
		if strings.Index(targetWord, v) == 0 {
			suffix, _ := strings.CutPrefix(targetWord, v)
			if memoCanConstruct(suffix, wordbank, memo) {
				memo[targetWord] = true
				return memo[targetWord]
			}
		}
	}

	memo[targetWord] = false
	return memo[targetWord]
}
