package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(dirs string) int {
	houses := 1
	const len = 2 * 8192
	x, y := len/2, len/2
	grid := [len][len]int{}
	grid[y][x] = 1

	for _, dir := range dirs {
		if dir == '^' {
			y++
		} else if dir == '>' {
			x++
		} else if dir == 'v' {
			y--
		} else if dir == '<' {
			x--
		}
		if grid[y][x] == 0 {
			houses++
			grid[y][x] = 1
		}
	}

	return houses
}

func part2(dirs string) int {
	houses := 1
	const len = 2 * 8192
	sx, sy, rx, ry := len/2, len/2, len/2, len/2
	grid := [len][len]int{}
	grid[sy][sx] = 1

	for i, dir := range dirs {
		if dir == '^' {
			if i%2 == 0 {
				sy++
			} else {
				ry++
			}
		} else if dir == '>' {
			if i%2 == 0 {
				sx++
			} else {
				rx++
			}
		} else if dir == 'v' {
			if i%2 == 0 {
				sy--
			} else {
				ry--
			}
		} else if dir == '<' {
			if i%2 == 0 {
				sx--
			} else {
				rx--
			}
		}
		if i%2 == 0 && grid[sy][sx] == 0 {
			houses++
			grid[sy][sx] = 1
		} else if i%2 == 1 && grid[ry][rx] == 0 {
			houses++
			grid[ry][rx] = 1
		}
	}

	return houses
}

func main() {
	args := os.Args[1:]
	fileName, part := args[0], args[1]

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	dirs := scanner.Text()

	if part == "1" {
		fmt.Println(part1(dirs))
	} else if part == "2" {
		fmt.Println(part2(dirs))
	}
}
