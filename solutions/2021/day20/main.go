package main

import (
	"bufio"
	"fmt"
	"os"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			if val == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func addPadding(board [][]int) [][]int {
	padded := make([][]int, len(board)+2)
	for y := range padded {
		row := make([]int, len(board[0])+2)
		if y != 0 && y != len(padded)-1 {
			for x := 0; x < len(board[0]); x++ {
				row[x+1] = board[y-1][x]
			}
		}
		padded[y] = row
	}
	return padded
}

func updateBoard(board [][]int, line string) [][]int {
	updated := make([][]int, len(board))
	for i := range board {
		updated[i] = make([]int, len(board[i]))
		copy(updated[i], board[i])
	}
	for y := 1; y < len(board)-1; y++ {
		for x := 1; x < len(board[0])-1; x++ {
			val := 0
			for i := y - 1; i <= y+1; i++ {
				for j := x - 1; j <= x+1; j++ {
					val = 2*val + board[i][j]
				}
			}
			if line[val] == '#' {
				updated[y][x] = 1
			} else {
				updated[y][x] = 0
			}
		}
	}
	return updated
}

func part1(strs []string) int {
	board := [][]int{}
	for _, str := range strs[2:] {
		row := []int{}
		for _, c := range str {
			if c == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		board = append(board, row)
	}
	line := strs[0]
	fmt.Println(line)
	fmt.Println()
	printBoard(board)

	padded := addPadding(addPadding(board))
	printBoard(padded)
	updated := updateBoard(padded, line)
	printBoard(updated)

	padded = addPadding(updated)
	printBoard(padded)
	updated = updateBoard(padded, line)
	printBoard(updated)

	ans := 0
	for y := range updated {
		for x := range updated[y] {
			if updated[y][x] == 1 {
				ans++
			}
		}
	}

	return ans
}

func part2(strs []string) int {
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
