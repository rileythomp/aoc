package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rileythomp/aoc/src/aoc"
	_ "github.com/rileythomp/aoc/src/aoc"
)

func IsHexColor(c string) bool {
	if len(c) != 7 || string(c[0]) != "#" {
		return false
	}
	for _, hex := range c[1:] {
		if !strings.Contains("0123456789abcdef", string(hex)) {
			return false
		}
	}
	return true
}

func hasRequired(list []string, required []string) bool {
	for _, req := range required {
		if !aoc.ContainsStr(list, req) {
			return false
		}
	}
	return true
}

func part1(strs []string) int {
	valid := 0
	curPassport := []string{}
	for _, line := range strs {
		if line == "" {
			if len(curPassport) > 6 && hasRequired(curPassport, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}) {
				valid++
			}
			curPassport = []string{}
			continue
		}
		parts := strings.Fields(line)
		for _, part := range parts {
			idpart := strings.Split(part, ":")
			curPassport = append(curPassport, idpart[0])
		}
	}
	if len(curPassport) > 6 && hasRequired(curPassport, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}) {
		valid++
	}
	return valid
}

func passportHasField(passport [][2]string, req string) string {
	for i := range passport {
		if req == passport[i][0] {
			return passport[i][1]
		}
	}
	return ""
}

func isDigits(str string) bool {
	for _, c := range str {
		if !strings.Contains("0123456789", string(c)) {
			return false
		}
	}
	return true
}

func isValid(passport [][2]string) bool {
	if len(passport) < 7 {
		return false
	}
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, req := range required {
		val := passportHasField(passport, req)
		if val == "" {
			return false
		} else if req == "byr" && (val < "1920" || val > "2002") {
			return false
		} else if req == "iyr" && (val < "2010" || val > "2020") {
			return false
		} else if req == "eyr" && (val < "2020" || val > "2030") {
			return false
		} else if req == "hgt" {
			if string(val[len(val)-2]) == "cm" && (len(val) != 5 || val[0:3] < "150" || val[0:3] > "193") {
				return false
			} else if string(val[len(val)-2]) == "in" && (len(val) != 4 || val[0:2] < "59" || val[0:2] > "76") {
				return false
			}
		} else if req == "hcl" && !IsHexColor(val) {
			return false
		} else if req == "ecl" && !aoc.ContainsStr([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, val) {
			return false
		} else if req == "pid" && (!isDigits(val) || len(val) != 9) {
			return false
		}
	}
	return true
}

func part2(strs []string) int {
	valid := 0
	curPassport := [][2]string{}
	for _, line := range strs {
		if line == "" {
			if isValid(curPassport) {
				valid++
			}
			curPassport = [][2]string{}
			continue
		}
		parts := strings.Fields(line)
		for _, part := range parts {
			idpart := strings.Split(part, ":")
			curPassport = append(curPassport, [2]string{idpart[0], idpart[1]})
		}
	}
	if isValid(curPassport) {
		valid++
	}
	return valid
}

func main() {
	level, name := getArgs()
	if level == "" || name == "" {
		return
	}

	file, _ := os.Open(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	strs := []string{}
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	if level == "1" {
		fmt.Println(part1(strs))
	} else if level == "2" {
		fmt.Println(part2(strs))
	}
}

func getArgs() (string, string) {
	args := os.Args[1:]
	var (
		level = "1"
		name  = "input.txt"
	)
	for i, arg := range args {
		if i == 0 && (level == "1" || level == "2") {
			level = arg
		} else if i == 0 {
			fmt.Printf("Level must be 1 or 2, got %s\n", arg)
			return "", ""
		} else if i == 1 {
			name = arg
		}
	}
	return level, name
}
