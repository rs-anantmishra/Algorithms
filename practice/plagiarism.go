/*
	Find the longest common subsequence, given
	string1: i have this book i like to read
	string2: she has this book i like to watch
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "i have this book i like to read"
	wordbank := "she has this book i like to watch"

	result := longestSubsequence(strings.Split(text, " "), strings.Split(wordbank, " "))
	fmt.Println(resultSubsequence, result)
}

var subsequence string
var resultSubsequence string

func longestSubsequence(text []string, wordbank []string) string {

	space := " "
	var result string
	//base case
	if len(text) == 0 {
		return result
	}

	for i, value := range text {
		for k, word := range wordbank {
			if value == word {
				subsequence = word + space + subsequence
				if len(resultSubsequence) == 0 || len(resultSubsequence) < len(subsequence) {
					resultSubsequence = subsequence
				}
				if longestSubsequence(text[i:], wordbank[k:]) == "" {
					return result
				}
			}
		}
	}

	return result
}

//<space> should be ignored to be used as a suffix
//(i have) this book i like to read
//for the same suffix and within a same contiguous sequence
////////subtract from the wordbank if a word is used.

//string1: (i have) this book i like to read
//string2: she has this book i like to watch
//this is a test sample from an interview
