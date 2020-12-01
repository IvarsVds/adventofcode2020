package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	result1, result2, err := run()
	if err != nil {
		panic(err)
	}

	fmt.Println("The result for first task is", result1)
	fmt.Println("The result for second task is", result2)
}

func run() (int, int, error) {
	input, err := getInput()
	if err != nil {
		return 0, 0, err
	}

	result1 := calcFirstTask(input)
	result2 := calcSecondTask(input)

	return result1, result2, nil
}

func getInput() ([]int, error) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return make([]int, 0), err
	}

	strSlice := strings.Split(string(input), "\n")
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		intSlice[i], _ = strconv.Atoi(str)
	}

	return intSlice, nil
}

// calcFirstTask finds two numbers that sum to 2020, then multiplies these numbers and returns the result
func calcFirstTask(input []int) int {
	var a, b int

	for i, num := range input {
		for _, num2 := range input[i+1:] {
			if num+num2 != 2020 {
				continue
			}

			a = num
			b = num2
		}
	}

	return a * b
}

func calcSecondTask(input []int) int {
	var a, b, c int

	for i, num := range input {
		for _, num2 := range input[i+1:] {
			for _, num3 := range input[i+2:] {
				if num+num2+num3 != 2020 {
					continue
				}

				a = num
				b = num2
				c = num3
			}
		}
	}

	return a * b * c
}
