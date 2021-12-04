package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRow(boardline string) []int {
	nums := strings.Split(boardline, " ")
	rets := []int{}
	for _, num := range nums {
		if num == "" {
			continue
		}
		ret, _ := strconv.Atoi(num)
		rets = append(rets, ret)
	}
	return rets
}

func isWinner(board [][]int) bool {
	for i, row := range board {
		horizWin, vertWin := true, true
		for j, val := range row {
			if val > -1 {
				horizWin = false
			}
			if board[j][i] > -1 {
				vertWin = false
			}
		}
		if horizWin || vertWin {
			return true
		}
	}
	return false
}

func boardScore(board [][]int) int {
	score := 0
	for _, row := range board {
		for _, val := range row {
			if val > -1 {
				score += val
			}
		}
	}
	return score
}

func getBoardAndCalls(strs []string) ([][][]int, []int) {
	input := strs[0]
	nums := strings.Split(input, ",")
	calls := []int{}
	for _, c := range nums {
		num, _ := strconv.Atoi(c)
		calls = append(calls, num)
	}
	boardlines := strs[1:]
	boards := [][][]int{}
	curBoard := [][]int{}
	counter := 0
	for _, boardline := range boardlines {
		if boardline != "" {
			row := getRow(boardline)
			curBoard = append(curBoard, row)
			counter = (counter + 1) % 5
			if counter == 0 {
				boards = append(boards, curBoard)
				curBoard = [][]int{}
			}
		}
	}
	return boards, calls
}

func part1(strs []string) int {
	boards, numsCalled := getBoardAndCalls(strs)
	for _, call := range numsCalled {
		for b, board := range boards {
			for i, row := range board {
				for j, val := range row {
					if val == call {
						boards[b][i][j] = -1
					}
					if isWinner(board) {
						return boardScore(board) * call
					}
				}
			}
		}
	}
	return 0
}

func contains(nums []int, num int) bool {
	for i := range nums {
		if nums[i] == num {
			return true
		}
	}
	return false
}

func part2(strs []string) int {
	boards, numsCalled := getBoardAndCalls(strs)
	lastScore := 0
	winners := []int{}
	for _, call := range numsCalled {
		for b, board := range boards {
			if contains(winners, b) {
				continue
			}
			for i, row := range board {
				nextboard := false
				for j, val := range row {
					if val == call {
						boards[b][i][j] = -1
					}
					if isWinner(board) {
						winners = append(winners, b)
						lastScore = boardScore(board) * call
						nextboard = true
						break
					}
				}
				if nextboard {
					break
				}
			}
		}
	}
	return lastScore
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
