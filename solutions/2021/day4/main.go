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
		// ret, _ := strconv.Atoi(num)
		if num == "" || num == " " || num == "\n" {
			continue
		}
		num = strings.Replace(num, " ", "", -1)
		num = strings.Replace(num, "\n", "", -1)
		num = strings.Replace(num, "\t", "", -1)
		ret, _ := strconv.Atoi(num)
		rets = append(rets, ret)
	}
	return rets
}

func isWinner(board [][]int) bool {
	// check rows
	for _, row := range board {
		allNeg := true
		for _, val := range row {
			if val > -1 {
				allNeg = false
			}
		}
		if allNeg {
			return true
		}
	}
	// check cols
	// for i := 0; i < 6; i++ {
	// 	allNeg := true
	// 	for j := 0; j < 5; j++ {
	// 		if board[j][i] > -1 {
	// 			allNeg = false
	// 		}
	// 	}
	// }
	for i, col := range board {
		allNeg := true
		for j := range col {
			if board[j][i] > -1 {
				allNeg = false
			}
		}
		if allNeg {
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

func part1(strs []string) int {
	calls := strs[0]
	_ = calls
	numsCalled := []int{}
	callnums := strings.Split(calls, ",")
	for _, c := range callnums {
		x, _ := strconv.Atoi(c)
		numsCalled = append(numsCalled, x)
	}

	boardlines := strs[1:]
	counter := 0
	boards := [][][]int{}
	curBoard := [][]int{}
	for _, boardline := range boardlines {
		if boardline != "" {
			// fmt.Println(boardline)
			row := getRow(boardline)
			curBoard = append(curBoard, row)
			counter = (counter + 1) % 5
			if counter == 0 {
				boards = append(boards, curBoard)
				curBoard = [][]int{}
			}
		}
	}

	for _, call := range numsCalled {
		for b, board := range boards {
			for i, row := range board {
				for j, val := range row {
					if val == call {
						boards[b][i][j] = (boards[b][i][j] + 1) * -1
					}
					// check board, return board score * call if winner
					if isWinner(board) {
						return boardScore(board) * call
					}
				}
			}
		}
	}

	return 0
}

func part2(strs []string) int {
	calls := strs[0]
	_ = calls
	numsCalled := []int{}
	callnums := strings.Split(calls, ",")
	for _, c := range callnums {
		x, _ := strconv.Atoi(c)
		numsCalled = append(numsCalled, x)
	}

	boardlines := strs[1:]
	counter := 0
	boards := [][][]int{}
	curBoard := [][]int{}
	for _, boardline := range boardlines {
		if boardline != "" {
			// fmt.Println(boardline)
			row := getRow(boardline)
			curBoard = append(curBoard, row)
			counter = (counter + 1) % 5
			if counter == 0 {
				boards = append(boards, curBoard)
				curBoard = [][]int{}
			}
		}
	}

	// have boards here
	lastScore := 0
	winners := []int{}
	for _, call := range numsCalled {
		for b, board := range boards {
			// skip if b in winners
			skip := false
			for _, w := range winners {
				if b == w {
					skip = true
				}
			}
			if skip {
				continue
			}
			nextboard := false
			for i, row := range board {
				for j, val := range row {
					if val == call {
						boards[b][i][j] = (boards[b][i][j] + 1) * -1
					}
					// check board, return board score * call if winner
					if isWinner(board) {
						fmt.Printf("done board %d with %d\n", b, call)
						winners = append(winners, b)
						fmt.Println(winners)
						lastScore = boardScore(board) * call
						nextboard = true
						if len(winners) == len(boards) {
							fmt.Println(board)
							fmt.Printf("Winning score %d, %d, %d\n", boardScore(board)*call, boardScore(board), call)
						}
					}
					if nextboard {
						break
					}
				}
				if nextboard {
					break
				}
			}
		}
	}
	fmt.Println(winners)
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
