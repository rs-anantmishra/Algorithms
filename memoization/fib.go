package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}
