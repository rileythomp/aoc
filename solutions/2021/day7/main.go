package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func part1(strs []string) int {
	vals := strings.Split(strs[0], ",")
	nums := []int{}
	min, max := math.MaxInt32, 0
	for i := range vals {
		num, _ := strconv.Atoi(vals[i])
		nums = append(nums, num)
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	minCount := math.MaxInt32
	curCount := 0
	for i := min; i <= max; i++ {
		for j := range nums {
			curCount += abs(i - nums[j])
		}
		if curCount < minCount {
			minCount = curCount
		}
		curCount = 0
	}
	return minCount
}

func part2(strs []string) int {
	vals := strings.Split(strs[0], ",")
	nums := []int{}
	min, max := math.MaxInt32, 0
	for i := range vals {
		num, _ := strconv.Atoi(vals[i])
		nums = append(nums, num)
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	minCount := math.MaxInt32
	curCount := 0
	for i := min; i <= max; i++ {
		for j := range nums {
			dist := abs(i - nums[j])
			curCount += (dist * (dist + 1)) / 2
		}
		if curCount < minCount {
			minCount = curCount
		}
		curCount = 0
	}
	return minCount
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
