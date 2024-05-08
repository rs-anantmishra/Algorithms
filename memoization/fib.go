package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	var computation string
	if len(os.Args) == 3 {
		computation = os.Args[2]
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

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

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
