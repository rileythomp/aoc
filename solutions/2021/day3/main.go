package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getBits(strs []string) []int {
	bits := make([]int, len(strs[0]))
	for _, str := range strs {
		for i, bit := range str {
			if bit == '1' {
				bits[i]++
			} else {
				bits[i]--
			}
		}
	}
	return bits
}

func part1(strs []string) int {
	bits := getBits(strs)
	gamma, epsilon := 0, 0
	for _, bit := range bits {
		gamma *= 2
		epsilon *= 2
		if bit > 0 {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon
}

func filter(strs []string, index int, val byte) []string {
	rets := []string{}
	for _, str := range strs {
		if str[index] == val {
			rets = append(rets, str)
		}
	}
	return rets
}

func getString(strs []string, neg, pos byte) string {
	i := 0
	for len(strs) > 1 {
		bits := getBits(strs)
		if bits[i] < 0 {
			strs = filter(strs, i, neg)
		} else {
			strs = filter(strs, i, pos)
		}
		i++
	}
	return strs[0]
}

func part2(strs []string) int {
	ogr, csr := 0, 0
	ogrs, csrs := getString(strs, '0', '1'), getString(strs, '1', '0')
	for i := range ogrs {
		ogr *= 2
		csr *= 2
		if ogrs[i] == '1' {
			ogr++
		}
		if csrs[i] == '1' {
			csr++
		}
	}
	return ogr * csr
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
	} else {
		fmt.Println("Level must be 1 or 2")
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
