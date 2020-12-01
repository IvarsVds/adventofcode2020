package main

import "testing"

func TestCalcFirstTask(t *testing.T) {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	test := calcFirstTask(testInput)
	if test != 514579 {
		t.Errorf("Expected result to be 514579, but got %v", test)
	}
}

func TestCalcSecondTask(t *testing.T) {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	test := calcSecondTask(testInput)
	if test != 241861950 {
		t.Errorf("Expected result to be 241861950, but got %v", test)
	}
}
