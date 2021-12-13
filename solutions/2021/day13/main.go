package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func part1(strs []string) int {
	board := [2000][2000]int{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		if strings.Contains(str, "fold") {
			break
		}
		xy := strings.Split(str, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		board[y][x] = 1
	}
	for _, str := range strs {
		if !strings.Contains(str, "fold") {
			continue
		}
		dirval := strings.Split(strings.Fields(str)[2], "=")
		dir, val := dirval[0], dirval[1]
		num, _ := strconv.Atoi(val)
		for y, row := range board {
			for x := range row {
				if board[y][x] == 0 {
					continue
				}
				if dir == "y" && y > num {
					board[y][x] = 0
					if 2*num-y >= 0 {
						board[2*num-y][x] = 1
					}
				} else if dir == "x" && x > num {
					board[y][x] = 0
					if 2*num-x >= 0 {
						board[y][2*num-x] = 1
					}
				}
			}
		}
		break
	}
	dots := 0
	for _, row := range board {
		for _, val := range row {
			if val > 0 {
				dots++
			}
		}
	}
	return dots
}

func part2(strs []string) int {
	points := [][2]int{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		if strings.Contains(str, "fold") {
			break
		}
		xy := strings.Split(str, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		points = append(points, [2]int{x, y})
	}
	for _, str := range strs {
		if !strings.Contains(str, "fold") {
			continue
		}
		dirval := strings.Split(strings.Fields(str)[2], "=")
		dir, val := dirval[0], dirval[1]
		num, _ := strconv.Atoi(val)
		for i, p := range points {
			if dir == "y" && p[1] > num {
				points[i] = [2]int{p[0], 2*num - p[1]}
			} else if dir == "x" && p[0] > num {
				points[i] = [2]int{2*num - p[0], p[1]}
			}
		}
	}
	boardstr := ""
	for y := 0; y < 7; y++ {
		for x := 0; x < 159; x++ {
			broke := false
			for _, p := range points {
				if p[0] == x && p[1] == y {
					boardstr += "#"
					broke = true
					break
				}
			}
			if !broke {
				boardstr += " "
			}
			boardstr += " "
		}
		boardstr += "\n"
	}
	fmt.Println(boardstr)
	return 0
}

func part2v0(strs []string) int {
	board := [2000][2000]int{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		if strings.Contains(str, "fold") {
			break
		}
		xy := strings.Split(str, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		board[y][x] = 1
	}
	for _, str := range strs {
		if !strings.Contains(str, "fold") {
			continue
		}
		dirval := strings.Split(strings.Fields(str)[2], "=")
		dir, val := dirval[0], dirval[1]
		num, _ := strconv.Atoi(val)
		for y, row := range board {
			for x := range row {
				if board[y][x] == 0 {
					continue
				}
				if dir == "y" && y > num {
					board[y][x] = 0
					if 2*num-y >= 0 {
						board[2*num-y][x] = 1
					}
				} else if dir == "x" && x > num {
					board[y][x] = 0
					if 2*num-x >= 0 {
						board[y][2*num-x] = 1
					}
				}
			}
		}
	}
	boardstr := ""
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if board[y][x] > 0 {
				boardstr += "#"
			} else {
				boardstr += "."
			}
			boardstr += " "
		}
		boardstr += "\n"
	}
	fmt.Println(boardstr)
	return 0
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
