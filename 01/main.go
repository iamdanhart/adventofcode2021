package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	prev := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		digit, _ := strconv.Atoi(line)
		if prev != 0 && digit > prev {
			count++
		}
		prev = digit
	}
	fmt.Println("Part 1 solution: " + strconv.Itoa(count))
}
func partTwo() {
	scanner := makeScanner(false)
	count := 0
	// I know the input size is > 3, else this is super sloppy to do
	scanner.Scan()
	l1 := scanner.Text()
	w1, _ := strconv.Atoi(l1)
	scanner.Scan()
	l2 := scanner.Text()
	w2, _ := strconv.Atoi(l2)
	scanner.Scan()
	l3 := scanner.Text()
	w3, _ := strconv.Atoi(l3)
	for scanner.Scan() {
		line := scanner.Text()
		digit, _ := strconv.Atoi(line)
		oldSum := w1 + w2 + w3
		newSum := w2 + w3 + digit

		if w2 != 0 && w3 != 0 && oldSum < newSum {
			count++
		}
		w1 = w2
		w2 = w3
		w3 = digit
	}
	fmt.Println("Part 2 solution: " + strconv.Itoa(count))
}
