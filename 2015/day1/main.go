package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(brackets string) int {
	count := 0
	for _, bracket := range brackets {
		if bracket == '(' {
			count++
		} else if bracket == ')' {
			count--
		}
	}
	return count
}

func part2(brackets string) int {
	count := 0
	for i, bracket := range brackets {
		if bracket == '(' {
			count++
		} else if bracket == ')' {
			count--
			if count < 0 {
				return i + 1
			}
		}
	}
	return -1
}

func main() {
	args := os.Args[1:]
	fileName, part := args[0], args[1]

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	brackets := scanner.Text()

	if part == "1" {
		fmt.Println(part1(brackets))
	} else if part == "2" {
		fmt.Println(part2(brackets))
	}
}
