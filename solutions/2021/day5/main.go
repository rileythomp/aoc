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

func getPoints(str string) (int, int, int, int) {
	parts := strings.Fields(str)
	p1, p2 := parts[0], parts[2]
	p1parts := strings.Split(p1, ",")
	p2parts := strings.Split(p2, ",")
	p1x, p1y := p1parts[0], p1parts[1]
	p2x, p2y := p2parts[0], p2parts[1]
	x1, _ := strconv.Atoi(p1x)
	y1, _ := strconv.Atoi(p1y)
	x2, _ := strconv.Atoi(p2x)
	y2, _ := strconv.Atoi(p2y)
	return x1, y1, x2, y2
}

func part1(strs []string) int {
	board := [1000][1000]int{}
	for _, str := range strs {
		x1, y1, x2, y2 := getPoints(str)
		miny, minx := min(y1, y2), min(x1, x2)
		maxy, maxx := y1+y2-miny, x1+x2-minx
		if minx != maxx && miny != maxy {
			continue
		}
		for i := miny; i <= maxy; i++ {
			for j := minx; j <= maxx; j++ {
				board[i][j] += 1
			}
		}
	}
	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] > 1 {
				count++
			}
		}
	}
	return count
}

func part2(strs []string) int {
	board := [1000][1000]int{}
	for _, str := range strs {
		x1, y1, x2, y2 := getPoints(str)
		miny, minx := min(y1, y2), min(x1, x2)
		maxy, maxx := y1+y2-miny, x1+x2-minx
		if minx != maxx && miny != maxy {
			if x1 == minx && miny == y1 || x1 == maxx && y1 == maxy {
				for i := miny; i <= maxy; i++ {
					for j := minx; j <= maxx; j++ {
						if i-miny == j-minx {
							board[i][j] += 1
						}
					}
				}
			} else {
				for i := miny; i <= maxy; i++ {
					for j := maxx; j >= minx; j-- {
						if i-miny == maxx-j {
							board[i][j] += 1
						}
					}
				}
			}
			continue
		}
		for i := miny; i <= maxy; i++ {
			for j := minx; j <= maxx; j++ {
				board[i][j] += 1
			}
		}
	}
	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] > 1 {
				count++
			}
		}
	}
	return count
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
