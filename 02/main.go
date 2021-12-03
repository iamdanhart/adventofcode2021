package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	partOne()
	duration := time.Since(start)
	partTwo()
	duration2 := time.Since(start)
	fmt.Printf("p1: %s, p2: %s\n", duration, duration2-duration)
}
func makeScanner(test bool) *bufio.Scanner {
	var f *os.File
	if test {
		f, _ = os.Open("inputs/testinput")
	} else {
		f, _ = os.Open("inputs/input")
	}
	reader := bufio.NewReader(f)
	return bufio.NewScanner(reader)
}
func partOne() {
	scanner := makeScanner(false)
	horPos := 0
	depth := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		dir := lineSplit[0]
		mag := lineSplit[1]
		magVal, _ := strconv.Atoi(mag)
		switch dir {
			case "forward":
				horPos += magVal
		case "down":
			depth += magVal
		case "up":
			depth -= magVal
		}
	}
	fmt.Println("Part 1 solution: ", horPos * depth)
}
func partTwo() {
	scanner := makeScanner(false)
	horPos := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		dir := lineSplit[0]
		mag := lineSplit[1]
		magVal, _ := strconv.Atoi(mag)
		switch dir {
		case "forward":
			horPos += magVal
			depth += aim * magVal
		case "down":
			aim += magVal
		case "up":
			aim -= magVal
		}
	}
	fmt.Println("Part 2 solution: ", horPos * depth)
}
