package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	os.Args = append(os.Args, "howsum")

	if os.Args[1] == "fib" {
		//os.Args = append(os.Args, "fib")
		//os.Args = append(os.Args, 100)
		//os.Args = append(os.Args, "m")
		CallFib(os.Args)
	}

	if os.Args[1] == "grid" {
		// os.Args = append(os.Args, "grid")
		// os.Args = append(os.Args, "18")
		// os.Args = append(os.Args, "18")
		// os.Args = append(os.Args, "m")
		start := time.Now()
		CallGridTraveler(os.Args)
		elapsed := time.Since(start)
		fmt.Println("elapsed time is ", elapsed)

	}

	if os.Args[1] == "cansum" {
		CallCanSum(os.Args)
	}

	if os.Args[1] == "howsum" {
		CallHowSum(os.Args)
	}

}
