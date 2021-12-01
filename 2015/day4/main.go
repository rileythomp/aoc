package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func part1(str string) int {
	i := 1
	for true {
		hash := fmt.Sprintf("%x\n", md5.Sum([]byte(str+strconv.Itoa(i))))
		if hash[0:5] == "00000" {
			return i
		}
		i++
	}
	return 0
}

func part2(str string) int {
	i := 1
	for true {
		hash := fmt.Sprintf("%x\n", md5.Sum([]byte(str+strconv.Itoa(i))))
		if hash[0:6] == "000000" {
			return i
		}
		i++
	}
	return 0
}

func main() {
	args := os.Args[1:]
	str, part := args[0], args[1]

	// file, _ := os.Open(fileName)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)

	// scanner.Scan()

	if part == "1" {
		fmt.Println(part1(str))
	} else if part == "2" {
		fmt.Println(part2(str))
	}
}
