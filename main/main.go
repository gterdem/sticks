package main

import (
	"github.com/gterdem/sticks/solver"
)

func main() {
	// input := "(9:00-10:00) - (9:00-9:30)"
	input := "(9:00-10:00, 10:00-10:30) - (9:00-9:30)"
	// input := "(9:00-10:00) minus (9:00-9:30)" //test for this input
	solver.Solve(input)
}
