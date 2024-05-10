package main

import "fmt"

func CallFib() {

	n := 9
	result := fib(n)

	fmt.Println(result)
}

func fib(n int) int {

	fib := 0
	slcFib := make([]int, n+1)

	//seed values 0 & 1
	slcFib[1] = 1

	for i, _ := range slcFib {

		//skip seed values
		if i == 0 {
			continue
		}

		next := i + 1
		nextPlusOne := i + 2

		if next < len(slcFib) {
			slcFib[next] = slcFib[i] + slcFib[next]
		}

		if nextPlusOne < len(slcFib) {
			slcFib[nextPlusOne] = slcFib[i] + slcFib[nextPlusOne]
		}
	}

	fib = slcFib[n]
	return fib
}
