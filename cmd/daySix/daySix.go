package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day Six, memory clean up.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	sourceData := getSourceData(filename)
	fmt.Println(sourceData)
	cycles, diff := getCycleLength(sourceData)
	fmt.Printf("It took %v cycles to find a repeat. Loop size was %v", cycles, diff)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSourceData(filename string) []int {
	f, err := os.Open(filename)
	check(err)
	reader := bufio.NewReader(f)
	numberScanner := bufio.NewScanner(reader)
	numberScanner.Split(bufio.ScanWords)
	var source []int
	var number int
	for numberScanner.Scan() {
		number, _ = strconv.Atoi(numberScanner.Text())
		source = append(source, number)
	}

	return source
}

func getCycleLength(memory []int) (int, int) {
	var count int
	found := false
	size := len(memory)
	patterns := make(map[string]int)
	var diff int
	for !found {
		count++
		rebalance(memory, size)
		found, diff = checkRepeat(memory, patterns, size, count)

	}
	return count, diff
}

func rebalance(memory []int, size int) {
	sourceIndex := findLargestIndex(memory, size)
	banks := memory[sourceIndex]
	memory[sourceIndex] = 0
	for banks > 0 {
		sourceIndex++
		memory[sourceIndex%size]++
		banks--
	}
}

func checkRepeat(memory []int, patterns map[string]int, size int, count int) (bool, int) {
	sig := getSig(memory, size)
	_, ok := patterns[sig]
	found := true
	if !ok {
		patterns[sig] = count
		found = false
	}
	diff := count - patterns[sig]
	return found, diff
}

func getSig(memory []int, size int) string {
	var sig string
	var curChar string
	for i := 0; i < size; i++ {
		curChar = strconv.Itoa(memory[i])
		sig = sig + string(curChar) + ":"
	}
	return sig
}

func findLargestIndex(memory []int, size int) int {
	var max int
	var index int
	for i := 0; i < size; i++ {
		if memory[i] > max {
			index = i
			max = memory[i]
		}
	}
	return index
}
