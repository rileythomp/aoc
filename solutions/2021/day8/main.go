package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func hasVal(str string, val string) bool {
	for _, c := range val {
		if !strings.Contains(str, string(c)) {
			return false
		}
	}
	return true
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
			numstr[1] = str
			strnum[str] = 1
		} else if len(str) == 3 {
			numstr[7] = str
			strnum[str] = 7
		} else if len(str) == 4 {
			numstr[4] = str
			strnum[str] = 4
		} else if len(str) == 7 {
			numstr[8] = str
			strnum[str] = 8
		}
	}
	oneval, fourval := numstr[1], numstr[4]
	for _, str := range input {
		if len(str) == 6 {
			if !hasVal(str, oneval) {
				numstr[6] = str
				strnum[str] = 6
			} else if hasVal(str, fourval) {
				numstr[9] = str
				strnum[str] = 9
			} else {
				numstr[0] = str
				strnum[str] = 0
			}

		} else if len(str) == 5 {
			if hasVal(str, oneval) {
				numstr[3] = str
				strnum[str] = 3
			} else if valParts(str, fourval) == 3 {
				numstr[5] = str
				strnum[str] = 5
			} else {
				numstr[2] = str
				strnum[str] = 2
			}
		}
	}
	val := 0
	for _, str := range output {
		for k, v := range strnum {
			if len(k) == len(str) && hasVal(k, str) {
				val = 10*val + v
			}
		}
	}
	return val
}

func part2(strs []string) int {
	count := 0
	val := 0
	_ = count
	_ = val
	for i := range strs {
		val += calcVal(strs[i])
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
	nums := []int{}
	for scanner.Scan() {
		str := scanner.Text()
		strs = append(strs, str)
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
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
