package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rileythomp/aoc/src/aoc"
	_ "github.com/rileythomp/aoc/src/aoc"
)

func part1(strs []string) int {
	rolls := 0
	roll := 1
	pos1, pos2 := 6, 7
	score1, score2 := 0, 0
	for score1 < 1000 && score2 < 1000 {
		if rolls%2 == 0 {
			pos1 += (3*roll + 3) % 10
			if pos1 > 10 {
				pos1 -= 10
			}
			score1 += pos1
		} else {
			pos2 += (3*roll + 3) % 10
			if pos2 > 10 {
				pos2 -= 10
			}
			score2 += pos2
		}
		roll += 3
		rolls++
	}
	return aoc.Min(score1, score2) * 3 * rolls
}

type gamestate struct {
	score1, score2 int
	pos1, pos2     int
	numgames       int // number of games up to this state
}

var freqmap = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func part2(strs []string) int {
	rolls := 0
	wins1, wins2 := 0, 0
	gameStates := aoc.Queue{gamestate{0, 0, 4, 8, 1}}
	for !gameStates.IsEmpty() {
		cur := gameStates.Pop().(gamestate)
		for i := 3; i <= 9; i++ {
			if rolls%2 == 0 {
				// add gamestate to gameStates
				pos1 := cur.pos1
				pos1 += i
				if pos1 > 10 {
					pos1 -= 10
				}
				score := cur.score1 + pos1
				new := gamestate{pos1, cur.pos2, score, cur.score2, cur.numgames * freqmap[i]}
				if score >= 21 {
					wins1 += new.numgames
					// fmt.Println("p1 wins got", wins1, "games")
					continue
				}
				gameStates.Push(new)
			} else {
				pos2 := cur.pos2
				pos2 += i
				if pos2 > 10 {
					pos2 -= 10
				}
				score := cur.score1 + pos2
				new := gamestate{cur.pos1, pos2, cur.score1, score, cur.numgames * freqmap[i]}
				if score >= 21 {
					wins2 += new.numgames
					// fmt.Println("p2 wins got", wins2, "games")
					continue
				}
				gameStates.Push(new)
			}
		}
		rolls++
	}
	return aoc.Max(wins1, wins2)
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
