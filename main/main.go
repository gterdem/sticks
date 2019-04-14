package main

import (
	"github.com/gterdem/sticks/solver"
)

func main() {
	input := "(9:00-10:00) - (9:00-9:30)" // ok
	// input := "(9:00-10:00) - (9:00-10:00)" // missing empty array print
	// input := "(9:00-9:30) - (9:30-15:00)" // ok
	// input := "(9:00-9:30, 10:00-10:30) - (9:15-10:15)" // ok
	// input := "(9:00-11:00, 13:00-15:00) - (9:00-9:15, 10:00-10:15, 12:30-16:00)"
	// input := "(9:00-10:00) minus (9:00-9:30)" //test for this input
	solver.Solve(input)
}