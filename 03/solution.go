package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// run() what?
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	res1 := solveFirstTask(input)
	res2 := solveSecondTask(input, res1)

	fmt.Println("The result for first task is", res1)
	fmt.Println("The result for second task is", res2)
}

func getInput() ([]string, error) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return make([]string, 0), err
	}

	strSlice := strings.Split(string(input), "\n")

	return strSlice, nil
}

func solveFirstTask(input []string) int {
	result := 0
	i := 0
	mp := 0

	for _, str := range input {
		pStr := prepString(str, mp)
		for (len(pStr) - 1) <= i {
			mp++
			pStr = prepString(str, mp)
		}

		if string(pStr[i]) == "#" {
			result++
		}

		i = i + 3
	}

	return result
}

func prepString(input string, mp int) string {
	if mp == 0 {
		return input
	}

	var sb strings.Builder

	for i := 0; i < mp; i++ {
		sb.WriteString(input)
	}

	return sb.String()
}

func solveSecondTask(input []string, i int) int {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go countTrees(ch1, input, 1, 1)
	go countTrees(ch2, input, 5, 1)
	go countTrees(ch3, input, 7, 1)
	go countTrees(ch4, input, 1, 2)

	res1 := <-ch1
	res2 := <-ch2
	res3 := <-ch3
	res4 := <-ch4

	return i * res1 * res2 * res3 * res4
}

func countTrees(ch chan int, input []string, toRight int, down int) {
	result := 0
	i := 0
	mp := 0

	for j := 0; j < len(input); j = j + down {
		pStr := prepString(input[j], mp)
		for (len(pStr) - 1) <= i {
			mp++
			pStr = prepString(input[j], mp)
		}

		if string(pStr[i]) == "#" {
			result++
		}

		i = i + toRight
	}

	ch <- result
}
