package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	_ "github.com/rileythomp/aoc/src/aoc"
)

func hexToBin(hex string) string {
	htob := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	bin := ""
	for _, h := range hex {
		bin += htob[string(h)]
	}
	return bin
}

func binstrval(bin string) int {
	val := 0
	for i := 0; i < len(bin); i++ {
		num, _ := strconv.Atoi(string(bin[i]))
		val = 2*val + num
	}
	return val
}

func calcVal(vals []int, typeId string) int {
	ans := 0
	if typeId == "000" {
		for _, v := range vals {
			ans += v
		}
	} else if typeId == "001" {
		ans = 1
		for _, v := range vals {
			ans *= v
		}
	} else if typeId == "010" {
		ans = math.MaxInt32
		for _, v := range vals {
			if v < ans {
				ans = v
			}
		}
	} else if typeId == "011" {
		ans = math.MinInt32
		for _, v := range vals {
			if v > ans {
				ans = v
			}
		}
	} else if typeId == "101" {
		if vals[0] > vals[1] {
			ans = 1
		}
	} else if typeId == "110" {
		if vals[0] < vals[1] {
			ans = 1
		}
	} else if typeId == "111" {
		if vals[0] == vals[1] {
			ans = 1
		}
	}
	return ans
}

func parsePacket(bin string) (int, int, int, bool) {
	if len(bin) < 11 {
		return 0, 0, 0, false
	}
	versionSum := binstrval(bin[0:3])
	val := 0
	pktBits := 0
	typeId := bin[3:6]
	if typeId == "100" {
		binstr := ""
		for i := 6; i < len(bin); i += 5 {
			binstr += bin[(i + 1):(i + 5)]
			if bin[i] == '0' {
				pktBits = i + 5
				break
			}
		}
		val = binstrval(binstr)
	} else if bin[6] == '0' {
		len := binstrval(bin[7:22])
		pktBits = 22
		vals := []int{}
		for pktBits-22 < len {
			ver, subval, bits, ok := parsePacket(bin[pktBits:])
			if !ok {
				break
			}
			versionSum += ver
			pktBits += bits
			vals = append(vals, subval)
		}
		val = calcVal(vals, typeId)
	} else if bin[6] == '1' {
		pkts := binstrval(bin[7:18])
		pktBits = 18
		vals := []int{}
		for i := 0; i < pkts; i++ {
			ver, subval, bits, ok := parsePacket(bin[pktBits:])
			if !ok {
				break
			}
			versionSum += ver
			pktBits += bits
			vals = append(vals, subval)
		}
		val = calcVal(vals, typeId)
	}
	return versionSum, val, pktBits, true
}

func part1(strs []string) int {
	hex := strs[0]
	binary := hexToBin(hex)
	versionSum, _, _, _ := parsePacket(binary)
	return versionSum
}

func part2(strs []string) int {
	hex := strs[0]
	binary := hexToBin(hex)
	_, val, _, _ := parsePacket(binary)
	return val
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
