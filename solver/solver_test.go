package solver

import (
	"testing"
)

//TestSample1 "(9:00-10:00) - (9:00-9:30) = (9:30-10:00)"
func TestSample1(t *testing.T) {
	input := "(9:00-10:00) - (9:00-9:30)"
	want := "(9:30-10:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestSample1_1 "(9:00-10:00) minus (9:00-9:30) = (9:30-10:00)"
func TestSample1_1(t *testing.T) {
	input := "(9:00-10:00) minus (9:00-9:30)"
	want := "(9:30-10:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestSample2 "(9:00-10:00) - (9:00-10:00) = ( )"
func TestSample2(t *testing.T) {
	input := "(9:00-10:00) - (9:00-10:00)"
	want := "()"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestSample3 "(9:00-9:30) - (9:30-15:00) = (9:00-9:30)"
func TestSample3(t *testing.T) {
	input := "(9:00-9:30) - (9:30-15:00)"
	want := "(9:00-9:30)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestSample4 "(9:00-9:30, 10:00-10:30) - (9:15-10:15) = (9:00-9:15, 10:15-10:30)"
func TestSample4(t *testing.T) {
	input := "(9:00-9:30, 10:00-10:30) - (9:15-10:15)"
	want := "(9:00-9:15, 10:15-10:30)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestSample5 "(9:00-11:00, 13:00-15:00) - (9:00-9:15, 10:00-10:15, 12:30-16:00) = (9:15-10:00, 10:15-11:00)"
func TestSample5(t *testing.T) {
	input := "(9:00-11:00, 13:00-15:00) - (9:00-9:15, 10:00-10:15, 12:30-16:00)"
	want := "(9:15-10:00, 10:15-11:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra1 "(00:00-06:45) minus (00:00-00:15) = (00:15-6:45)"
func TestExtra1(t *testing.T) {
	input := "(00:00-06:45) minus (00:00-00:15)"
	want := "(00:15-6:45)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra2 "(00:00-08:15) minus (08:15-12:00) = (00:00-08:15)"
func TestExtra2(t *testing.T) {
	input := "(00:00-08:15) minus (08:15-12:00)"
	want := "(00:00-8:15)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra3 "(1:00-3:00) minus (1:00-4:00) = ()"
func TestExtra3(t *testing.T) {
	input := "(1:00-3:00) minus (1:00-4:00)"
	want := "()"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra4 "(00:00-3:00, 7:00-8:15) minus (2:00-7:30) = (00:00-2:00, 7:30-8:15)"
func TestExtra4(t *testing.T) {
	input := "(00:00-3:00, 7:00-8:15) minus (2:00-7:30)"
	want := "(00:00-2:00, 7:30-8:15)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra5 "(1:00-5:00, 7:00-12:00) minus (0:00-6:00, 7:15-8:00, 13:00-15:00) = (7:00-7:15, 8:00-12:00)"
func TestExtra5(t *testing.T) {
	input := "(1:00-5:00, 7:00-12:00) minus (0:00-6:00, 7:15-8:00, 13:00-15:00)"
	want := "(7:00-7:15, 8:00-12:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra6 "(23:00-3:00) minus (23:40-1:00) = (23:00-23:40, 1:00-3:00)"
func TestExtra6(t *testing.T) {
	input := "(23:00-3:00) minus (23:40-1:00)"
	want := "(23:00-23:40, 1:00-3:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra7 "(00:00-3:00) minus (00:00-02:00) = (02:00-03:00)"
func TestExtra7(t *testing.T) {
	input := "(00:00-3:00) minus (00:00-02:00)"
	want := "(2:00-3:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra8 "(00:00-3:00) minus (02:00-03:00) = (00:00-02:00)"
func TestExtra8(t *testing.T) {
	input := "(00:00-3:00) minus (02:00-03:00)"
	want := "(00:00-2:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}

//TestExtra9 "(01:00-03:00, 05:00-07:00) minus (02:00-06:00) =(1:00-2:00, 6:00-7:00)"
func TestExtra9(t *testing.T) {
	input := "(01:00-03:00, 05:00-07:00) minus (02:00-06:00)"
	want := "(1:00-2:00, 6:00-7:00)"
	resultStr := StringifyResult(Solve(input))
	if resultStr != want {
		t.Errorf("Result incorrect, got: %v, wanted: %v", resultStr, want)
	}
}
