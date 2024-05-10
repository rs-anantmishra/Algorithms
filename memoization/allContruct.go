package main

import (
	"fmt"
	"strings"
)

func CallAllConstruct(args []string) {

	fmt.Println("abcdef", allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))
	fmt.Println("purple", allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))

}

// doesnt work fully
func allConstruct(targetWord string, wordbank []string) []string {

	var result []string
	//base cases
	if targetWord == "" {
		return []string{}
	}

	for _, v := range wordbank {
		if strings.Index(targetWord, v) == 0 {
			suffix, _ := strings.CutPrefix(targetWord, v)
			suffixWays := allConstruct(suffix, wordbank)

			if len(suffixWays) == 0 {
				suffixWays = append(suffixWays, v)
			}

			fmt.Println(suffixWays)
			//for _, val := range suffixWays {
			//}

			//result = append(result, targetWays...)
		}
	}

	return result
}
