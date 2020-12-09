package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	res1 := solveFirstTask(input)

	fmt.Println(res1)
}

func getInput() ([][]string, error) {
	results := make([][]string, 0)

	f, err := os.Open("./input.txt")
	if err != nil {
		return results, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	group := make([]string, 0)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			results = append(results, group)
			group = make([]string, 0)
		}

		group = append(group, scanner.Text())
	}
	results = append(results, group)

	if err := scanner.Err(); err != nil {
		return make([][]string, 0), err
	}

	return results, nil
}

func solveFirstTask(input [][]string) int {
	result := 0
	for _, ss := range input {
		str := concatString(ss)
		result = result + countUniqueChars(str)
	}

	return result
}

func countUniqueChars(s string) int {
	count := 0
	m := make(map[rune]bool)

	// loop over string
	for _, ltr := range s {
		// see if rune is already present in map, if false
		// add it and increase count
		if _, value := m[ltr]; !value {
			m[ltr] = true
			count++
		}
	}

	return count
}

func concatString(input []string) string {
	var sb strings.Builder

	for _, str := range input {
		sb.WriteString(str)
	}

	return sb.String()
}
