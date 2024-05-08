package main

import (
	"fmt"
	"strconv"
)

// Complexity: Space - O(2^n), Time - O(n)
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

// Complexity: Space - O(n), Time - O(n)
func fibMemo(n int, m map[int]int) int {

	_, ok := m[n]
	if ok {
		return m[n]
	}
	if n <= 2 {
		return 1
	}

	m[n] = fibMemo(n-2, m) + fibMemo(n-1, m)
	return m[n]

}

// considering fib's starting position 0 with value 0
// 0	1	1	2	3	5	8	13  21
// 0	1	2	3	4	5	6	7	8

func CallFib(args []string) {
	var computation string
	n, _ := strconv.Atoi(args[2])
	if len(args) == 4 {
		computation = args[3]
	}
	var m = make(map[int]int)

	switch {
	case computation == "memo" || computation == "Memo" || computation == "m":
		fmt.Println("memoized", fibMemo(n, m))
	case n < 40:
		fmt.Println("standard", fib(n))
	default:
		fmt.Println("memoized", fibMemo(n, m))
	}
}
