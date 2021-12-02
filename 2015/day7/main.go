package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wire struct {
	Val      uint16
	Gate     string
	Resolved bool
}

func processGate(gate string, wires map[string]Wire) (uint16, bool) {
	// NOT X
	if strings.Contains(gate, "NOT") {
		gateParts := strings.Split(gate, " ")
		input := gateParts[1]
		if num, err := strconv.Atoi(input); err == nil {
			return ^uint16(num), true
		}
		wire, ok := wires[input]
		if ok && wire.Resolved {
			return ^wire.Val, true
		}
	}
	return 0, false
}

func part1(strs []string) int {
	wires := map[string]Wire{}
	for _, str := range strs {
		parts := strings.Split(str, " -> ")
		gate, wire := parts[0], parts[1]
		val, ok := processGate(gate, wires)
		if ok {
			wires[wire] = Wire{
				Val:      val,
				Gate:     gate,
				Resolved: true,
			}
		} else {
			wires[wire] = Wire{
				Gate:     gate,
				Resolved: false,
			}
		}
	}
	a, _ := wires["a"]
	return int(a.Val)
}

func part2(strs []string) int {
	return 0
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
