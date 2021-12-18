package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func addSnailfish(sf1, sf2 string) string {
	return "[" + sf1 + "," + sf2 + "]"
}

func getMagnitude(sf string, start, end int) int {
	pair := sf[start+1 : end]
	vals := strings.Split(pair, ",")
	a, _ := strconv.Atoi(vals[0])
	b, _ := strconv.Atoi(vals[1])
	return 3*a + 2*b
}

func calcMagnitude(sf string) int {
	if len(strings.Split(sf, ",")) == 2 {
		return getMagnitude(sf, 0, len(sf)-1)
	}
	open := 0
	for i, c := range sf {
		if c == '[' {
			open = i
		}
		if c == ']' {
			sf = sf[:open] + fmt.Sprint(getMagnitude(sf, open, i)) + sf[(i+1):]
			break
		}
	}
	return calcMagnitude(sf)
}

func canExplode(sf string) (int, bool) {
	opens := 0
	for i, c := range sf {
		if c == '[' {
			opens++
		} else if c == ']' {
			opens--
		}
		if opens > 4 {
			return i, true
		}
	}
	return -1, false
}

func isDigit(c string) bool {
	return strings.Contains("0123456789", c)
}

func canSplit(sf string) (int, bool) {
	for i, c := range sf {
		if c == '[' || c == ',' {
			if isDigit(string(sf[i+1])) && isDigit(string(sf[i+2])) {
				return i, true
			}
		}
	}
	return -1, false
}

func explodeSnailfish(sf string, start int) string {
	end := 0
	for i := range sf[start+1:] {
		if sf[start+1:][i] == ']' {
			end = i
			break
		}
	}
	pair := sf[start+1 : start+1+end]
	vals := strings.Split(pair, ",")
	left, _ := strconv.Atoi(vals[0])
	right, _ := strconv.Atoi(vals[1])

	end = start + end + 1
	_, _ = left, right
	firsthalf := sf[:start]
	lasthalf := sf[end+1:]

	lnumend, lnumstart := -1, -1
	for i := start - 1; i >= 0; i-- {
		if isDigit(string(sf[i])) {
			lnumend = i
			break
		}
	}
	if lnumend >= 0 {
		lnumstart = lnumend
		for i := lnumend - 1; i >= 0; i-- {
			if isDigit(string(sf[i])) {
				lnumstart = i
			} else {
				break
			}
		}
		leftadd, _ := strconv.Atoi(sf[lnumstart : lnumend+1])
		leftadd += left
		firsthalf = sf[:lnumstart] + fmt.Sprint(leftadd) + sf[lnumend+1:start]
	}

	rnumstart, rnumend := -1, -1
	for i := end + 1; i < len(sf); i++ {
		if isDigit(string(sf[i])) {
			rnumstart = i
			break
		}
	}
	if rnumstart >= 0 {
		rnumend = rnumstart
		for i := rnumstart + 1; i <= len(sf); i++ {
			if isDigit(string(sf[i])) {
				rnumend = i
			} else {
				break
			}
		}
		rightadd, _ := strconv.Atoi(sf[rnumstart : rnumend+1])
		rightadd += right
		lasthalf = sf[end+1:rnumstart] + fmt.Sprint(rightadd) + sf[rnumend+1:]
	}
	return firsthalf + "0" + lasthalf
}

func splitSnailfish(sf string, start int) string {
	val := ""
	end := 0
	for i := range sf[start+1:] {
		if isDigit(string(sf[start+1:][i])) {
			val += string(sf[start+1:][i])
		} else if sf[start+1:][i] == ',' || sf[start+1:][i] == ']' {
			end = i
			break
		}
	}
	num, _ := strconv.Atoi(val)
	left, right := num/2, num/2
	if num%2 != 0 {
		right++
	}
	pair := fmt.Sprintf("[%d,%d]", left, right)
	return sf[:start+1] + pair + sf[end+start+len(val)-1:]
}

func part1(strs []string) int {
	ans := strs[0]
	strs = strs[1:]
	for _, str := range strs {
		ans = addSnailfish(ans, str)
		for {
			if index, ok := canExplode(ans); ok {
				ans = explodeSnailfish(ans, index)
			} else if index, ok := canSplit(ans); ok {
				ans = splitSnailfish(ans, index)
			} else {
				break
			}
		}
	}
	return calcMagnitude(ans)
}

func part2(strs []string) int {
	ans := 0
	for _, str1 := range strs {
		for _, str2 := range strs {
			if str1 == str2 {
				continue
			}
			magnitude := part1([]string{str1, str2})
			if magnitude > ans {
				ans = magnitude
			}
		}
	}
	return ans
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
