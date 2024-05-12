package main

import (
	"fmt"
	"strings"
)

func CallCountConstruct() {

	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))

}

func countConstruct(word string, dictionary []string) int {

	result := 0

	//handle offset by 1 nature
	table := make([]int, len(word)+1)

	//seed all values as false
	for i := range table {
		table[i] = 0
	}

	//seed 0th as True (you can use 0 words to create an empty string)
	table[0] = 1

	for k, v := range table {
		if v > 0 {
			for _, val := range dictionary {
				if strings.Index(word, val) == 0 {
					fwdIndex := len(val)
					if len(table) > (k + fwdIndex) {
						table[k+fwdIndex] = table[k+fwdIndex] + 1
					}

				}
			}
		}
		if len(word) > 0 {
			word = word[1:]
		}
	}

	result = table[6]
	return result
}
