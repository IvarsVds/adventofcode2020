package main

import "testing"

func TestSolveFirstTask(t *testing.T) {
	testInput := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	test := solveFirstTask(testInput)
	if test != 7 {
		t.Errorf("Expected result to be 7, but got %v", test)
	}
}

func TestSolveSecondTask(t *testing.T) {
	testInput := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	test := solveSecondTask(testInput, 7)
	if test != 336 {
		t.Errorf("Expected result to be 336, but got %v", test)
	}
}
