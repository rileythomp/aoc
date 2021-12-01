package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(strs []string) int {
	return 0
}

func part2(strs []string) int {
	return 0
}

func main() {
	args := os.Args[1:]
	fileName, part := args[0], args[1]

	file, _ := os.Open(fileName)
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

	if part == "1" {
		fmt.Println(part1(strs))
	} else if part == "2" {
		fmt.Println(part2(strs))
	}
}
