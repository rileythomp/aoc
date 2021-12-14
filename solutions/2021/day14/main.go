package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func part1(strs []string) int {
	polymer := strs[0]
	rules := map[string]string{}
	for _, str := range strs[2:] {
		parts := strings.Fields(str)
		rules[parts[0]] = parts[2]
	}
	for step := 0; step < 10; step++ {
		curpoly := polymer
		for i := 0; i < len(curpoly)-1; i++ {
			pair := string(curpoly[i]) + string(curpoly[i+1])
			insert := rules[pair]
			polymer = polymer[:(2*i+1)] + insert + polymer[(2*i+1):]
		}
	}
	countmap := map[string]int{}
	for _, c := range polymer {
		if _, ok := countmap[string(c)]; ok {
			countmap[string(c)]++
		} else {
			countmap[string(c)] = 0
		}
	}
	max, min := math.MinInt32, math.MaxInt32
	for _, v := range countmap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}

func part2(strs []string) int {
	polymer := strs[0]
	rules := map[string]string{}
	for _, str := range strs[2:] {
		parts := strings.Fields(str)
		rules[parts[0]] = parts[2]
	}
	for step := 0; step < 40; step++ {
		curpoly := polymer
		for i := 0; i < len(curpoly)-1; i++ {
			pair := string(curpoly[i]) + string(curpoly[i+1])
			insert := rules[pair]
			polymer = polymer[:(2*i+1)] + insert + polymer[(2*i+1):]
		}
	}
	countmap := map[string]int{}
	for _, c := range polymer {
		if _, ok := countmap[string(c)]; ok {
			countmap[string(c)]++
		} else {
			countmap[string(c)] = 0
		}
	}
	max, min := math.MinInt32, math.MaxInt32
	for _, v := range countmap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
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
