package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	checkSum, divSum := getChecksum(filename)
	fmt.Printf("Checksum value: %v\tdivSum: %v\n", checkSum, divSum)
}

func getChecksum(filename string) (int, int) {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var sum int
	var divSum int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lineReader := bufio.NewReader(strings.NewReader(line))
		wordScanner := bufio.NewScanner(lineReader)
		wordScanner.Split(bufio.ScanWords)
		lineValue, divValue := processLine(wordScanner)
		sum += lineValue
		divSum += divValue
	}
	return sum, divSum
}

func processLine(wordScanner *bufio.Scanner) (int, int) {
	var min int
	var max int
	wordScanner.Scan()
	current, _ := strconv.Atoi(wordScanner.Text())
	min = current
	max = current
	var entries []int
	entries = append(entries, current)
	for wordScanner.Scan() {
		current, _ = strconv.Atoi(wordScanner.Text())
		entries = append(entries, current)
		if current > max {
			max = current
		}
		if current < min {
			min = current
		}
	}
	difference := max - min
	var divResult int
	sort.Sort(sort.Reverse(sort.IntSlice(entries)))
	lineSize := len(entries)
	for x := 0; x < lineSize-1; x++ {
		for y := x + 1; y < lineSize; y++ {
			if entries[x]%entries[y] == 0 {

				divResult = entries[x] / entries[y]
				fmt.Printf("X:%v\tY:%v\tResult:%v\n", entries[x], entries[y], divResult)
				x = lineSize + 1
				y = lineSize + 1

			}
		}
	}

	return difference, divResult
}
