package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(strs []string) int {
	lights := [1000][1000]int{}
	lightsOn := 0
	for _, str := range strs {
		parts := strings.Split(str, " ")
		action := parts[1]
		if action != "off" && action != "on" {
			action = "toggle"
		}

		topLeft := parts[len(parts)-3]
		tlCoords := strings.Split(topLeft, ",")
		tlX, _ := strconv.Atoi(tlCoords[0])
		tlY, _ := strconv.Atoi(tlCoords[1])

		bottomRight := parts[len(parts)-1]
		brCoords := strings.Split(bottomRight, ",")
		brX, _ := strconv.Atoi(brCoords[0])
		brY, _ := strconv.Atoi(brCoords[1])

		for y := tlY; y <= brY; y++ {
			for x := tlX; x <= brX; x++ {
				if action == "on" {
					if lights[y][x] == 0 {
						lightsOn++
					}
					lights[y][x] = 1
				} else if action == "off" {
					if lights[y][x] == 1 {
						lightsOn--
					}
					lights[y][x] = 0
				} else if action == "toggle" {
					if lights[y][x] == 0 {
						lightsOn++
					} else if lights[y][x] == 1 {
						lightsOn--
					}
					lights[y][x] = 1 - lights[y][x]
				}
			}
		}
	}
	return lightsOn
}

func part2(strs []string) int {
	lights := [1000][1000]int{}
	brightness := 0
	for _, str := range strs {
		parts := strings.Split(str, " ")
		action := parts[1]
		if action != "off" && action != "on" {
			action = "toggle"
		}

		topLeft := parts[len(parts)-3]
		tlCoords := strings.Split(topLeft, ",")
		tlX, _ := strconv.Atoi(tlCoords[0])
		tlY, _ := strconv.Atoi(tlCoords[1])

		bottomRight := parts[len(parts)-1]
		brCoords := strings.Split(bottomRight, ",")
		brX, _ := strconv.Atoi(brCoords[0])
		brY, _ := strconv.Atoi(brCoords[1])

		for y := tlY; y <= brY; y++ {
			for x := tlX; x <= brX; x++ {
				if action == "on" {
					brightness++
					lights[y][x] = lights[y][x] + 1
				} else if action == "off" {
					if lights[y][x] > 0 {
						brightness--
						lights[y][x] = lights[y][x] - 1
					}
				} else if action == "toggle" {
					brightness += 2
					lights[y][x] = lights[y][x] + 2
				}
			}
		}
	}
	return brightness
}

func main() {
	level, fileName := getArgs()
	if level == "" || fileName == "" {
		return
	}

	file, _ := os.Open(fileName)
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
		level    = "1"
		fileName = "input.txt"
	)
	for i, arg := range args {
		if i == 0 && (level == "1" || level == "2") {
			level = arg
		} else if i == 0 {
			fmt.Printf("Level must be 1 or 2, got %s\n", arg)
			return "", ""
		} else if i == 1 {
			fileName = arg
		}
	}
	return level, fileName
}
