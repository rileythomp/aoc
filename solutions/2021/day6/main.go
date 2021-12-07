package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(strs []string) int {
	line := strs[0]
	vals := strings.Split(line, ",")
	nums := []int{}
	for i := range vals {
		num, _ := strconv.Atoi(vals[i])
		nums = append(nums, num)
	}
	for i := 0; i < 80; i++ {
		for j := range nums {
			nums[j]--
			if nums[j] < 0 {
				nums[j] = 6
				nums = append(nums, 8)
			}
		}
	}
	return len(nums)
}

func part2(strs []string) int {
	line := strs[0]
	vals := strings.Split(line, ",")
	days := make([]int, 9)
	for i := range vals {
		d, _ := strconv.Atoi(vals[i])
		days[d]++
	}
	for i := 0; i < 256; i++ {
		days0 := days[0]
		days[0], days[1], days[2] = days[1], days[2], days[3]
		days[3], days[4], days[5] = days[4], days[5], days[6]
		days[6], days[7], days[8] = days0+days[7], days[8], days0
	}
	return days[0] + days[1] + days[2] + days[3] + days[4] + days[5] + days[6] + days[7] + days[8]
}

func part2_old(strs []string) int {
	// very slow, 3:30 to calculate answer
	line := strs[0]
	vals := strings.Split(line, ",")
	d1, d2 := 1, 256
	stack := make([][2]int, math.MaxInt32)
	length, children := 0, 0
	for i := range vals {
		num, _ := strconv.Atoi(vals[i])
		stack[length] = [2]int{d1, num - 1}
		length++
		children++
	}
	cache := make(map[int]int)
	curParent, curChildren := 0, 0
	for length > 0 {
		daynum := stack[length-1]
		length--
		d1, n := daynum[0], daynum[1]
		if d1+n >= d2 {
			continue
		}
		if n != 8 {
			if val, ok := cache[n+1]; ok {
				children += val
				continue
			}
			cache[curParent] = curChildren
			curChildren = 0
			curParent = n + 1
		}
		kids := int(math.Ceil(float64(d2-d1-n) / 7.0))
		curChildren += kids
		children += kids
		for d := d1 + n + 1; d <= d2; d += 7 {
			stack[length] = [2]int{d, 8}
			length++
		}
	}
	return children
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
