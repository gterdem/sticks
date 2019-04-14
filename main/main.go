package main

import (
	"fmt"
	"github.com/gterdem/sticks/solver"
)

func main() {
	// input := "(9:00-10:00) - (9:00-9:30)" // ok
	// input := "(9:00-10:00) - (9:00-10:00)" // ok
	// input := "(9:00-9:30) - (9:30-15:00)" // ok
	// input := "(9:00-9:30, 10:00-10:30) - (9:15-10:15)" // ok
	// input := "(9:00-11:00, 13:00-15:00) - (9:00-9:15, 10:00-10:15, 12:30-16:00)" // ok
	// input := "(9:00-10:00) minus (9:00-9:30)" // ok
	// input := "(01:00-03:00, 05:00-07:00) - (02:00-06:00)" // extra boundry test
	input := "(00:20-01:00, 02:00-03:00, 04:00-05:00) - (00:30-02:30, 02:45-04:30)" // extra boundry test-2
	results := solver.Solve(input)
	resultStr := solver.StringifyResult(results)
	fmt.Println(resultStr)
	// solver.PrintResults(results)
}
