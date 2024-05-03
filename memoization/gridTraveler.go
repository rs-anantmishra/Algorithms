package main

import (
	"fmt"
	"strconv"
)

type Grid struct {
	Row int
	Col int
}

func GridTravelerMemo(g Grid, m map[Grid]int) int {
	//fast-return case
	_, ok := m[g]
	if ok {
		return m[g]
	}
	//base cases
	if g.Row == 0 || g.Col == 0 {
		return 0
	}
	if g.Row == 1 || g.Col == 1 {
		return 1
	}

	m[g] = GridTravelerMemo(Grid{g.Row - 1, g.Col}, m) + GridTravelerMemo(Grid{g.Row, g.Col - 1}, m)
	return m[g]
}

func GridTraveler(m int, n int) int {

	//base cases
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return GridTraveler(m-1, n) + GridTraveler(m, n-1)
}

func CallGridTraveler(args []string) {

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
