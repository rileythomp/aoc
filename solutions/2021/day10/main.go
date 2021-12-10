package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var brackets = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func part1(strs []string) int {
	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	score := 0
	for _, str := range strs {
		stack := Stack{}
		for _, c := range str {
			char := string(c)
			if strings.Contains("({[<", char) {
				stack.Push(char)
			} else if char != brackets[stack.Pop()] {
				score += scores[char]
				break
			}
		}
	}
	return score
}

func part2(strs []string) int {
	scoreMap := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
	scores := []int{}
	for _, str := range strs {
		stack := Stack{}
		score, broke := 0, false
		for _, c := range str {
			char := string(c)
			if strings.Contains("({[<", char) {
				stack.Push(char)
			} else if char != brackets[stack.Pop()] {
				broke = true
				break
			}
		}
		if broke {
			continue
		}
		for !stack.IsEmpty() {
			score = 5*score + scoreMap[stack.Pop()]
		}
		scores = append(scores, score)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	return scores[len(scores)/2]
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
