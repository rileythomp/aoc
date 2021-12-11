package main

import (
	"bufio"
	"fmt"
	"os"

	aoc "github.com/rileythomp/aoc/src/aoc"
)

var adjacents = [][2]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Print(val)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func part1(strs []string) int {
	board := [][]int{}
	for _, str := range strs {
		nums := aoc.StrToNums(str)
		board = append(board, nums)
	}
	flashes := 0
	for step := 0; step < 100; step++ {
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				board[y][x]++
			}
		}
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				stack := Stack{{x, y}}
				for !stack.IsEmpty() {
					coord := stack.Pop()
					if board[coord[1]][coord[0]] > 9 {
						flashes++
						board[coord[1]][coord[0]] = -1
						for _, adj := range adjacents {
							if coord[1]+adj[1] >= 0 && coord[1]+adj[1] < len(board) &&
								coord[0]+adj[0] >= 0 && coord[0]+adj[0] < len(board) && board[coord[1]+adj[1]][coord[0]+adj[0]] >= 0 {
								board[coord[1]+adj[1]][coord[0]+adj[0]]++
								stack.Push([2]int{coord[0] + adj[0], coord[1] + adj[1]})
							}
						}
					}
				}
			}
		}
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				if board[y][x] < 0 {
					board[y][x] = 0
				}
			}
		}
	}
	return flashes
}

func part2(strs []string) int {
	board := [][]int{}
	for _, str := range strs {
		nums := aoc.StrToNums(str)
		board = append(board, nums)
	}
	flashes := 0
	step := 0
	for true {
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				board[y][x]++
			}
		}
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				stack := Stack{{x, y}}
				for !stack.IsEmpty() {
					coord := stack.Pop()
					if board[coord[1]][coord[0]] > 9 {
						flashes++
						board[coord[1]][coord[0]] = -1
						for _, adj := range adjacents {
							if coord[1]+adj[1] >= 0 && coord[1]+adj[1] < len(board) &&
								coord[0]+adj[0] >= 0 && coord[0]+adj[0] < len(board) && board[coord[1]+adj[1]][coord[0]+adj[0]] >= 0 {
								board[coord[1]+adj[1]][coord[0]+adj[0]]++
								stack.Push([2]int{coord[0] + adj[0], coord[1] + adj[1]})
							}
						}
					}
				}
			}
		}
		flashed := 0
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				if board[y][x] < 0 {
					flashed++
					board[y][x] = 0
				}
			}
		}
		step++
		if flashed == 100 {
			return step
		}
	}
	return -1
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

type Stack [][2]int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(e [2]int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() [2]int {
	if s.IsEmpty() {
		return [2]int{}
	}
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func (s *Stack) Top() [2]int {
	if s.IsEmpty() {
		return [2]int{}
	}
	e := (*s)[len(*s)-1]
	return e
}
