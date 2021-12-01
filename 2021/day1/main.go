package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)  

func part1(nums []int) int {
    increases := 0
    for i, num := range nums[1:] {
        if num > nums[i] {
            increases++
        }
    }
    return increases
}

func part2(nums []int) int {
    increases := 0
    for i, num := range nums[3:] {
        if num > nums[i] {
            increases++
        }
    }
    return increases
}

func main() {
    args := os.Args[1:]
    fileName, part := args[0], args[1]

    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var nums []int
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, num)
    }

    if part == "1" {
        fmt.Println(part1(nums))
    } else if part == "2" {
        fmt.Println(part2(nums))
    }
}

