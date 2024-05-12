package main

import (
	"fmt"
	"strings"
)

func CallCanConstruct() {

	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
}

func canConstruct(word string, dictionary []string) bool {

	result := false

	//handle offset by 1 nature
	table := make([]bool, len(word)+1)

	//seed all values as false
	for i := range table {
		table[i] = false
	}

	//seed 0th as True (you can use 0 words to create an empty string)
	table[0] = true

	for k, v := range table {
		if v {
			for _, val := range dictionary {
				if strings.Index(word, val) == 0 {
					fwdIndex := len(val)
					if len(table) > (k + fwdIndex) {
						table[k+fwdIndex] = true
					}

				}
			}
		}
		if len(word) > 0 {
			word = word[1:]
		}
	}

	return result
}
