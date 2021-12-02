package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(strs []string) int {
	h, d := 0, 0
	for _, str := range strs {
		parts := strings.Split(str, " ")
		num, _ := strconv.Atoi(parts[1])
		dir := parts[0]
		if dir == "forward" {
			h += num
		} else if dir == "down" {
			d += num
		} else if dir == "up" {
			d -= num
		}
	}
	return h * d
}

func part2(strs []string) int {
	h, d, aim := 0, 0, 0
	for _, str := range strs {
		parts := strings.Split(str, " ")
		num, _ := strconv.Atoi(parts[1])
		dir := parts[0]
		if dir == "forward" {
			h += num
			d += num * aim
		} else if dir == "down" {
			aim += num
		} else if dir == "up" {
			aim -= num
		}
	}
	return h * d
}

func main() {
	level, fileName := getArgs()
	if level == "" || fileName == "" {
		return
	}

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

	if level == "1" {
		fmt.Println(part1(strs))
	} else if level == "2" {
		fmt.Println(part2(strs))
	}
}

func getArgs() (string, string) {
	args := os.Args[1:]
	var (
		level    = "1"
		fileName = "input.txt"
	)
	for i, arg := range args {
		if i == 0 && (level == "1" || level == "2") {
			level = arg
		} else if i == 0 {
			fmt.Printf("Level must be 1 or 2, got %s\n", arg)
			return "", ""
		} else if i == 1 {
			fileName = arg
		}
	}
	return level, fileName
}
