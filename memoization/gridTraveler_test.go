package main

import (
	"testing"
)

func BenchmarkGridTraveler(b *testing.B) {
	var args []string
	args = append(args, "nil")
	args = append(args, "grid")
	args = append(args, "18")
	args = append(args, "18")
	args = append(args, "m")

	for i := 0; i < b.N; i++ {
		CallGridTraveler(args)
	}

}
