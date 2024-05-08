/*
	Find the longest common subsequence, given
	string1: i have this book i like to read
	string2: she has this book i like to watch

	As such this is not a dynamic programming problem
	since we are trying to find longest consecutive common
	subsequence - since its consecutive there are no
	repeating problems.
*/

package main

import (
	"fmt"
	"strings"
)

func Plagiarism() {
	space := " "
	resultSubsequence := []string{}
	result := ""

	text := "i have this book i like to read"
	wordbank := "she has this book i like to watch"

	for len(text) > 0 {

		if result == "" {
			text, _ = strings.CutPrefix(text, strings.Split(text, " ")[0])
		} else {
			text, _ = strings.CutPrefix(text, result)
		}
		text = strings.Trim(text, space)
		result = longestSubsequence(text, strings.Split(wordbank, " "))

		resultSubsequence = append(resultSubsequence, result)
	}

	fmt.Println(resultSubsequence)
}

func longestSubsequence(text string, sequence []string) string {

	var subsequence string
	space := " "
	//base case
	if len(text) == 0 {
		return subsequence
	}

	for i, word := range sequence {
		if strings.Index(text, word) == 0 {

			subsequence = subsequence + word
			suffix, _ := strings.CutPrefix(text, (word + space))
			if longestSubsequence(suffix, sequence[i:]) == "" {
				return subsequence
			}
		}
	}

	return subsequence
}
