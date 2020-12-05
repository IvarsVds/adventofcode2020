package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := getInput("./input.txt")
	if err != nil {
		panic(err)
	}

	res1 := solveFirstTask(input)
	res2 := solveSecondTask(input)

	fmt.Println("The result for first task is", res1)
	fmt.Println("The result for second task is", res2)
}

// read input line by line && concat strings
func getInput(fname string) ([]string, error) {
	results := make([]string, 0)

	f, err := os.Open(fname)
	if err != nil {
		return make([]string, 0), err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	str := ""
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			results = append(results, str)
			str = ""
		}

		str = str + scanner.Text() + " "
	}
	// add last string
	results = append(results, str)

	if err := scanner.Err(); err != nil {
		return make([]string, 0), err
	}

	return results, nil
}

func validatePassport(data string, strict bool) bool {
	byr, _ := regexp.MatchString("byr:.+\\s?", data)
	eyr, _ := regexp.MatchString("eyr:.+\\s?", data)
	iyr, _ := regexp.MatchString("iyr:.+\\s?", data)
	ecl, _ := regexp.MatchString("ecl:.+\\s?", data)
	pid, _ := regexp.MatchString("pid:.+\\s?", data)
	hcl, _ := regexp.MatchString("hcl:.+\\s?", data)
	cid, _ := regexp.MatchString("cid:.+\\s?", data)
	hgt, _ := regexp.MatchString("hgt:.+\\s?", data)

	if !strict {
		return byr && eyr && iyr && ecl && pid && hcl && hgt
	}

	return cid && byr && eyr && iyr && ecl && pid && hcl && hgt
}

func solveFirstTask(input []string) int {
	result := 0

	for _, str := range input {
		valid := validatePassport(str, false)
		if !valid {
			continue
		}
		result++
	}

	return result
}

// Second Task
// DOESN'T WORK

func validatePassport2(s string, strict bool) bool {
	byrs := findString("byr:\\d{4}\\s*", s)
	if byrs == "" {
		return false
	}
	byr, _ := strconv.Atoi(findString("\\d{4}", byrs))
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyrs := findString("iyr:\\d{4}\\s*", s)
	if iyrs == "" {
		return false
	}
	iyr, _ := strconv.Atoi(findString("\\d{4}", iyrs))
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyrs := findString("eyr:\\d{4}\\s*", s)
	if eyrs == "" {
		return false
	}
	eyr, _ := strconv.Atoi(findString("\\d{4}", eyrs))
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt := findString("hgt:\\d+(cm|in)", s)
	if hgt == "" {
		return false
	}
	hgtsys := findString("(cm|in)", hgt)
	hgtval, _ := strconv.Atoi(findString("\\d+", hgt))
	if hgtsys == "cm" && hgtval < 150 && hgtval > 193 {
		return false
	} else if hgtsys == "in" && hgtval < 59 && hgtval > 76 {
		return false
	}

	hcl := findString("hcl:#[0-9a-f]{6}\\s*", s)
	if hcl == "" {
		return false
	}

	ecl := findString("ecl:(amb|blu|brn|gry|grn|hzl|oth)\\s*", s)
	if ecl == "" {
		return false
	}

	pid := findString("pid:\\d{9}\\s*", s)
	if pid == "" {
		return false
	}

	if strict {
		// check cid here
	}

	return true
}

func findString(rp string, s string) string {
	re := regexp.MustCompile(rp)
	return re.FindString(s)
}

func solveSecondTask(input []string) int {
	result := 0

	for _, str := range input {
		valid := validatePassport2(str, false)
		if !valid {
			continue
		}
		result++
	}

	return result
}
