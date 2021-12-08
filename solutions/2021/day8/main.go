package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(strs []string) int {
	count := 0
	for _, line := range strs {
		parts := strings.Split(line, "|")
		output := strings.Fields(parts[1])
		for _, val := range output {
			if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
				count++
			}
		}
	}
	return count
}

func valParts(str, val string) int {
	parts := 0
	for _, c := range val {
		if strings.Contains(str, string(c)) {
			parts++
		}
	}
	return parts
}

func calcVal(line string) int {
	parts := strings.Split(line, "|")
	input, output := strings.Fields(parts[0]), strings.Fields(parts[1])
	numstr, strnum := make(map[int]string), make(map[string]int)
	for _, str := range input {
		if len(str) == 2 {
			numstr[1], strnum[str] = str, 1
		} else if len(str) == 3 {
			numstr[7], strnum[str] = str, 7
		} else if len(str) == 4 {
			numstr[4], strnum[str] = str, 4
		} else if len(str) == 7 {
			numstr[8], strnum[str] = str, 8
		}
	}
	oneval, fourval := numstr[1], numstr[4]
	for _, str := range input {
		if len(str) == 6 {
			if valParts(str, oneval) != len(oneval) {
				numstr[6], strnum[str] = str, 6
			} else if valParts(str, fourval) == len(fourval) {
				numstr[9], strnum[str] = str, 9
			} else {
				numstr[0], strnum[str] = str, 0
			}

		} else if len(str) == 5 {
			if valParts(str, oneval) == len(oneval) {
				numstr[3], strnum[str] = str, 3
			} else if valParts(str, fourval) == 3 {
				numstr[5], strnum[str] = str, 5
			} else {
				numstr[2], strnum[str] = str, 2
			}
		}
	}
	val := 0
	for _, str := range output {
		for k, v := range strnum {
			if len(k) == len(str) && valParts(k, str) == len(str) {
				val = 10*val + v
			}
		}
	}
	return val
}

func part2(strs []string) int {
	val := 0
	for _, str := range strs {
		val += calcVal(str)
	}
	return val
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
		str := scanner.Text()
		strs = append(strs, str)
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

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	i := len(*s) - 1
	str := (*s)[i]
	*s = (*s)[:i]
	return str
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	str := (*s)[len(*s)-1]
	return str
}
