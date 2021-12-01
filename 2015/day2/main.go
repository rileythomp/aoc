package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(boxes []string) int {
	area := 0
	for _, box := range boxes {
		dimensions := strings.Split(box, "x")
		h, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		l, _ := strconv.Atoi(dimensions[2])
		area += (2*(h*w+h*l+w*l) + min(h*w, min(h*l, w*l)))
	}
	return area
}

func part2(boxes []string) int {
	feet := 0
	for _, box := range boxes {
		dimensions := strings.Split(box, "x")
		h, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		l, _ := strconv.Atoi(dimensions[2])
		feet += (2*(h+w+l-max(h, max(w, l))) + h*w*l)
	}
	return feet
}

func main() {
	args := os.Args[1:]
	fileName, part := args[0], args[1]

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var boxes []string
	for scanner.Scan() {
		box := scanner.Text()
		boxes = append(boxes, box)
	}

	if part == "1" {
		fmt.Println(part1(boxes))
	} else if part == "2" {
		fmt.Println(part2(boxes))
	}
}
