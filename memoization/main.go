package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if os.Args[1] == "fib" {
		//os.Args = append(os.Args, "fib")
		//os.Args = append(os.Args, 100)
		//os.Args = append(os.Args, "m")
		callFib(os.Args)
	}

	if os.Args[1] == "grid" {
		// os.Args = append(os.Args, "grid")
		// os.Args = append(os.Args, "18")
		// os.Args = append(os.Args, "18")
		// os.Args = append(os.Args, "m")
		callGridTraveler(os.Args)
	}

}

func callGridTraveler(args []string) {

	var computation string
	//In how many ways can you travel to the goal of a grid of dim=m*n
	//only move down or right. //start: top-left //end: bottom-right
	if len(args) == 5 {
		computation = args[4]
	}

	m, _ := strconv.Atoi(args[2])
	n, _ := strconv.Atoi(args[3])

	var g Grid
	g.Row = m
	g.Col = n
	var mapster = make(map[Grid]int)

	switch {
	case computation == "memo" || computation == "Memo" || computation == "m":
		fmt.Println(GridTravelerMemo(g, mapster))
	default:
		fmt.Println(GridTraveler(m, n))
	}

}

func callFib(args []string) {
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
