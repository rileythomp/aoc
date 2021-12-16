package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rileythomp/aoc/src/aoc"
	_ "github.com/rileythomp/aoc/src/aoc"
)

type P struct{ x, y int }

var nbrs = []P{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func getMinPath(board [][]int) int {
	minBoard := make([][]int, len(board))
	for i := range board {
		minBoard[i] = make([]int, len(board[i]))
		copy(minBoard[i], board[i])
	}
	visited := map[P]bool{{0, 0}: true}
	queue := aoc.Queue{P{0, 0}}
	for !queue.IsEmpty() {
		p := queue.Pop().(P)
		for _, nbr := range nbrs {
			nbrx, nbry := p.x+nbr.x, p.y+nbr.y
			newp := P{nbrx, nbry}
			if nbrx < 0 || nbrx >= len(board[0]) || nbry < 0 || nbry >= len(board) {
				continue
			}
			if ok := visited[newp]; ok {
				if board[nbry][nbrx]+minBoard[p.y][p.x] < minBoard[nbry][nbrx] {
					minBoard[nbry][nbrx] = board[nbry][nbrx] + minBoard[p.y][p.x]
					queue.Push(newp)
				}
			} else {
				minBoard[nbry][nbrx] = board[nbry][nbrx] + minBoard[p.y][p.x]
				queue.Push(newp)
			}
			visited[newp] = true
		}
	}
	return minBoard[len(minBoard)-1][len(minBoard[0])-1] - 1
}

func getBoard(strs []string) [][]int {
	board := [][]int{}
	for _, str := range strs {
		nums := aoc.StrToNums(str)
		board = append(board, nums)
	}
	return board
}

func part1(strs []string) int {
	board := getBoard(strs)
	return getMinPath(board)
}

func part2(strs []string) int {
	inboard := getBoard(strs)
	board := make([][]int, 5*len(inboard))
	for i := range board {
		board[i] = make([]int, 5*len(inboard[0]))
	}
	for y, row := range board {
		for x := range row {
			relx, rely := x%len(inboard[0]), y%len(inboard)
			boardx, boardy := x/len(inboard[0]), y/len(inboard)
			board[y][x] = inboard[rely][relx] + boardx + boardy
			if board[y][x] > 9 {
				board[y][x] -= 9
			}
		}
	}
	return getMinPath(board)
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
