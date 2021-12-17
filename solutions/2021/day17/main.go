package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func getBounds(strs []string) (int, int, int, int) {
	input := strs[0]
	parts := strings.Split(input, " ")
	xparts := strings.Split(parts[2], "..")
	yparts := strings.Split(parts[3], "..")
	xstart, xend := xparts[0][2:5], xparts[1][0:3]
	ystart, yend := yparts[0][2:6], yparts[1][0:3]
	x1, _ := strconv.Atoi(xstart)
	x2, _ := strconv.Atoi(xend)
	y1, _ := strconv.Atoi(ystart)
	y2, _ := strconv.Atoi(yend)
	return x1, x2, y1, y2
}

func part1(strs []string) int {
	x1, x2, y1, y2 := getBounds(strs)
	ans := 0
	for x := 0; x <= x2; x++ {
		for y := y1; y <= -1*y1; y++ {
			vx, vy := x, y
			curx, cury := 0, 0
			maxHeight := 0
			for {
				curx += vx
				cury += vy
				if cury > maxHeight {
					maxHeight = cury
				}
				if x1 <= curx && curx <= x2 && y1 <= cury && cury <= y2 {
					if maxHeight > ans {
						ans = maxHeight
					}
					break
				} else if (vx == 0 && (curx < x1 || curx > x2)) || (cury < y1) {
					break
				}
				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				vy--
			}
		}
	}
	return ans
}

func part2(strs []string) int {
	x1, x2, y1, y2 := getBounds(strs)
	ans := 0
	for x := 0; x <= x2; x++ {
		for y := y1; y <= -1*y1; y++ {
			vx, vy := x, y
			curx, cury := 0, 0
			for {
				curx += vx
				cury += vy
				if x1 <= curx && curx <= x2 && y1 <= cury && cury <= y2 {
					ans++
					break
				} else if (vx == 0 && (curx < x1 || curx > x2)) || (cury < y1) {
					break
				}
				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				vy--
			}
		}
	}
	return ans
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
