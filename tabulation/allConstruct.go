package main

import (
	"fmt"
	"strings"
)

func CallAllConstruct() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef"}))
}

func allConstruct(word string, dictionary []string) [][]string {

	var table [][][]string

	for k := 0; k < len(word)+1; k++ {
		table = append(table, [][]string{})
	}

	//seed initial value
	table[0] = [][]string{[]string{}}

	for k, v := range table {
		if len(v) >= 0 {
			for _, val := range dictionary {
				if strings.Index(word, val) == 0 {
					fwdIndex := len(val)
					if len(table) > (k + fwdIndex) {
						for _, prefix := range v {
							if len(prefix) == 0 {
								table[k+fwdIndex] = append(table[k+fwdIndex], []string{val})
							} else {
								prefix = append(prefix, val)
								table[k+fwdIndex] = append(table[k+fwdIndex], prefix)
							}
						}
					}
				}
			}
		}
		if len(word) > 0 {
			word = word[1:]
		}
	}

	return table[6]
}
