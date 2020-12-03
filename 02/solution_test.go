package main

import "testing"

func TestCalcFirstTask(t *testing.T) {
	testInput := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	test := calcFirstTask(testInput)
	if test != 2 {
		t.Errorf("Expected result to be 2, but got %v", test)
	}
}

func TestCalcSecondTask(t *testing.T) {
	testInput := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	test := calcSecondTask(testInput)
	if test != 1 {
		t.Errorf("Expected result to be 1, but got %v", test)
	}
}
