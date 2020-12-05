package main

import "testing"

func TestSolveFirstTask(t *testing.T) {
	testInput, _ := getInput("./testinput.txt")
	if len(testInput) != 4 {
		t.Errorf("Expected to have 4 input values, but got %v", len(testInput))
		t.Errorf("Output: %s", testInput)
	}

	test := solveFirstTask(testInput)
	if test != 2 {
		t.Errorf("Expected result to be 2, but got %v", test)
	}
}

func TestSolveSecondTask(t *testing.T) {
	testInput, _ := getInput("./testinput2.txt")

	test := solveSecondTask(testInput)
	if test != 4 {
		t.Errorf("Expected result to be 4, but got %v", test)
	}
}
