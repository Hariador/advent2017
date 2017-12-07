package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day two. Spreadsheet check sum.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	steps := buildSteps(filename)
	stepCount := runSteps(steps)
	fmt.Printf("It took %v steps to exit\n", stepCount)
}

func buildSteps(filename string) []int {
	var steps []int
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		steps = append(steps, value)
	}
	return steps
}

func runSteps(steps []int) int {
	var count int
	end := len(steps) - 1
	fmt.Printf("Processing %v steps\n", end)
	fmt.Println(steps)
	var ip int
	var next int
	exited := false
	for !exited {
		count++

		next = next + steps[ip]
		if steps[ip] >= 3 {
			steps[ip]--
		} else {
			steps[ip]++
		}
		//fmt.Printf("IP: %v\tNext: %v\tStep: %v\n", ip, next, steps[ip])
		if next < 0 || next > end {
			exited = true
		} else {
			ip = next
		}

	}
	return count
}
