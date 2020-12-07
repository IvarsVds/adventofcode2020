package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := getInput("./input.txt")
	if err != nil {
		panic(err)
	}

	res1, _ := solveFirstTask(input)

	fmt.Println(res1)
}

func getInput(fname string) ([]string, error) {
	input, err := ioutil.ReadFile(fname)
	if err != nil {
		return make([]string, 0), err
	}

	strSlice := strings.Split(string(input), "\n")

	return strSlice, nil
}

func createRows() []int {
	rows := make([]int, 128)

	for i := range rows {
		rows[i] = i
	}

	return rows
}

// return slice of ids for 2nd task
func solveFirstTask(input []string) (int, []int) {
	result := 0
	ids := make([]int, 0)

	for _, str := range input {
		rows := createRows()
		cols := []int{0, 1, 2, 3, 4, 5, 6, 7}

		for _, ltr := range str {
			if string(ltr) == "F" {
				ri := len(rows) / 2
				rows = rows[:ri]
			} else if string(ltr) == "B" {
				ri := len(rows) / 2
				rows = rows[ri:]
			} else if string(ltr) == "L" {
				ci := len(cols) / 2
				cols = cols[:ci]
			} else if string(ltr) == "R" {
				ci := len(cols) / 2
				cols = cols[ci:]
			}
		}

		id := rows[0]*8 + cols[0]
		if id > result {
			result = id
		}

		ids = append(ids, id)
	}

	return result, ids
}

func solveSecondTask(input []string) int {

	return 0
}
