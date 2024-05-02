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
