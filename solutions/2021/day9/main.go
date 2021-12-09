package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var pairs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func part1(strs []string) int {
	var (
		board = [][]int{}
		total = 0
	)
	for _, str := range strs {
		nums := []int{}
		for _, c := range str {
			nums = append(nums, int(c-'0'))
		}
		board = append(board, nums)
	}
	for y, row := range board {
		for x, val := range row {
			neighbours, lowerNeighbours := 0, 0
			for _, pair := range pairs {
				iy, ix := y+pair[0], x+pair[1]
				if iy >= 0 && iy < len(board) && ix >= 0 && ix < len(board[y]) {
					neighbours++
					if board[iy][ix] > val {
						lowerNeighbours++
					}
				}

			}
			if lowerNeighbours == neighbours {
				total += val + 1
			}
		}
	}
	return total
}

func part2(strs []string) int {
	board := [][]int{}
	for _, str := range strs {
		nums := []int{}
		for _, c := range str {
			nums = append(nums, int(c-'0'))
		}
		board = append(board, nums)
	}
	basinSizes := []int{}
	for y, row := range board {
		for x, val := range row {
			if val < 0 || val == 9 {
				continue
			}
			board[y][x] *= 9
			basinSize := 1
			stack := Stack{[]int{y, x}}
			for !stack.IsEmpty() {
				coord := stack.Pop()
				cy, cx := coord[0], coord[1]
				for _, pair := range pairs {
					iy, ix := cy+pair[0], cx+pair[1]
					if iy >= 0 && iy < len(board) && ix >= 0 && ix < len(row) && board[iy][ix] != 9 {
						board[iy][ix] = 9
						stack.Push([]int{iy, ix})
						basinSize++
					}

				}
			}
			basinSizes = append(basinSizes, basinSize)
		}
	}
	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
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

type Stack [][]int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str []int) {
	*s = append(*s, str)
}

func (s *Stack) Pop() []int {
	if s.IsEmpty() {
		return []int{}
	}
	i := len(*s) - 1
	str := (*s)[i]
	*s = (*s)[:i]
	return str
}

func (s *Stack) Top() []int {
	if s.IsEmpty() {
		return []int{}
	}
	str := (*s)[len(*s)-1]
	return str
}
