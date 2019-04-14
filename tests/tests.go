package tests

import (
	"github.com/gterdem/sticks/solver"
	"testing"
)

//TestSample1 "(9:00-10:00) - (9:00-9:30) = (9:30-10:00)"
func TestSample1(t *testing.T) {
	input := "(9:00-10:00) - (9:00-9:30)"
	want := "(9:30-10:00)"
	resultStr := solver.StringifyResult(solver.Solve(input))
	if resultStr != "(9:30-10:00)" {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}
