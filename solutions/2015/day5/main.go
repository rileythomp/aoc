package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasThreeVowels(str string) bool {
	vowels := 0
	for _, char := range str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels++
			if vowels == 3 {
				return true
			}
		}
	}
	return false
}

func hasDoubleLetter(str string) bool {
	for i := range str[1:] {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func noBadSubstrings(str string) bool {
	for i := range str[1:] {
		pair := string(str[i]) + string(str[i+1])
		if pair == "ab" || pair == "cd" || pair == "pq" || pair == "xy" {
			return false
		}
	}
	return true
}

func isNice(str string) bool {
	return hasThreeVowels(str) && hasDoubleLetter(str) && noBadSubstrings(str)
}

func part1(words []string) int {
	nices := 0
	for _, word := range words {
		if isNice(word) {
			nices++
		}
	}
	return nices
}

func hasPairRepeat(str string) bool {
	pairs := make(map[string]int)
	for i := range str[1:] {
		pair := string(str[i]) + string(str[i+1])
		if index, ok := pairs[pair]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			pairs[pair] = i
		}
	}
	return false
}

func hasSandwich(str string) bool {
	for i := range str[2:] {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func isNice2(str string) bool {
	return hasPairRepeat(str) && hasSandwich(str)
}

func part2(words []string) int {
	nices := 0
	for _, word := range words {
		if isNice2(word) {
			nices++
		}
	}
	return nices
}

func main() {
	args := os.Args[1:]
	fileName, part := args[0], args[1]

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	if part == "1" {
		fmt.Println(part1(words))
	} else if part == "2" {
		fmt.Println(part2(words))
	}
}
