package main

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
