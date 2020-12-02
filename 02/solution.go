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

	// note to self, come back to this and use channels
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

func getInput() ([]string, error) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return make([]string, 0), err
	}

	strSlice := strings.Split(string(input), "\n")

	return strSlice, nil
}

func calcFirstTask(input []string) int {
	result := 0

	for _, str := range input {
		isPassValid := validatePassword(str)
		if isPassValid {
			result++
		}
	}

	return result
}

func validatePassword(data string) bool {
	strSplit := strings.Split(data, " ")

	// 0 min 1 max
	minmax := strings.Split(strSplit[0], "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])

	letter := string(strSplit[1][0])
	pw := strSplit[2]

	// go's strings.Count is a bit dumb, hope this works without writing a loop
	count := strings.Count(pw, string(letter))

	return min <= count && count <= max
}

func calcSecondTask(input []string) int {
	result := 0

	for _, str := range input {
		isPassValid := validatePassword2(str)
		if isPassValid {
			result++
		}
	}

	return result
}

func validatePassword2(data string) bool {
	strSplit := strings.Split(data, " ")

	// 0 min 1 max
	postions := strings.Split(strSplit[0], "-")
	pos1, _ := strconv.Atoi(postions[0])
	pos2, _ := strconv.Atoi(postions[1])

	letter := string(strSplit[1][0])
	pw := strSplit[2]

	lp1 := string(pw[pos1-1])
	lp2 := string(pw[pos2-1])

	if lp1 == letter && lp2 != letter {
		return true
	} else if lp1 != letter && lp2 == letter {
		return true
	}

	return false
}
