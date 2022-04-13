package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/rileythomp/aoc/src/aoc"
)

x	y	z

x	y 	z
x   z  -y
x  -y  -z
x  -z   y

-x  y	z
-x  z  -y
-x -y  -z
-x -z   y

y   x   z
y   z  -x
y  -x  -z
y  -z   x

-y  x   z
-y  z  -x
-y -x  -z
-y -z   x

z   y   x
z   x  -y
z  -y  -x
z  -x   y

-z  y   x
-z  x  -y
-z -y  -x
-z -x   y

type coord struct {
	x, y, z int
}

func calcOverlaps(scanner0 []coord, scanner []coord) int {
	for _, coord := range scanner {

	}
	return 0
}

func determineLocation(scanner0 []coord, scanner []coord) {

}

func part1(strs []string) int {
	scanners := [][]coord{}
	curScanner := []coord{}
	for _, str := range strs {
		if len(str) == 0 {
			scanners = append(scanners, curScanner)
			curScanner = []coord{}
			continue
		}
		if strings.Contains(str, "scanner") {
			continue
		}
		coords := strings.Split(str, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		curScanner = append(curScanner, coord{x, y, z})
	}
	scanner0 := scanners[0]
	for _, scanner := range scanners[1:] {
		overlaps := calcOverlaps(scanner0, scanner)
		if overlaps >= 12 {
			determineLocation(scanner0, scanner)
		}
	}
	return 0
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
