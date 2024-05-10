/*
this is a dynamic programming problem since we are not asking for consecutive sequence.
example: text1=abcde, text2=ace // Here answer will be ace.
See how it is not consecutive but ordering is still preserved.
*/

package main

import (
	"fmt"
)

var text1 string
var text2 string
var matchSeq string
var lcs string

func CallLCS() {

	text1 = "AGGTAB"
	text2 = "GXTXAYB"

	result := calculateLCS(text1, text2, len(text1), len(text1), "")
	fmt.Println(result)

	resultCount := calculateLCS2()
	fmt.Println(resultCount)
}

// func calculateLCS(text string, seq string, lenText int, lenSeq int) string {
func calculateLCS(text string, seq string, lenSuffix int, iterations int, subsequence string) string {

	//base cases
	if lenSuffix == 0 {
		if iterations > 0 {
			iterations--
			if len(lcs) == 0 || len(lcs) < len(subsequence) {
				lcs = subsequence
			}
			return calculateLCS(text1[len(text1)-iterations:], text2, iterations, iterations, "")
		} else {
			return lcs
		}
	}
	if len(text) == 0 && len(seq) == 0 {
		iterations--
		if len(lcs) == 0 || len(lcs) < len(subsequence) {
			lcs = subsequence
		}
		return calculateLCS(text1[len(text1)-iterations:], text2, iterations, iterations, "")
	}

	if len(text) == 0 || len(seq) == 0 {
		lenSuffix--
		return calculateLCS(text[len(text)-lenSuffix:], matchSeq, lenSuffix, iterations, subsequence)
	}

	if text[0] == seq[0] {
		matchSeq = seq[1:]
		subsequence = subsequence + string(text[0])
		return calculateLCS(text[1:], seq[1:], len(text[1:]), iterations, subsequence)
	}

	return calculateLCS(text, seq[1:], lenSuffix, iterations, subsequence)
}

func calculateLCS2() int {

	return 0
}
