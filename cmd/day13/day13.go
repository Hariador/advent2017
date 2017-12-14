package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day Eleven.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("dayEleven <filename>")
	}

	filename := os.Args[1]
	fw := getSourceData(filename)

	fmt.Printf("Took a delay of %v picoseconds\n", fw)
}

func getSourceData(filename string) int {
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	var remaing map[int]bool
	remaing = make(map[int]bool)
	for y := 0; y <= 5000000; y++ {
		remaing[y] = true

	}
	fmt.Println("Initialized")
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println("*****************")
		line := scanner.Text()
		lineValues := strings.Split(line, ":")
		pos, _ := strconv.Atoi(lineValues[0])
		width, err := strconv.Atoi(strings.Trim(lineValues[1], " "))
		check(err)
		fmt.Printf("POS: %v\tDept: %v\n", pos, width)
		removeBases(remaing, pos, width)
		fmt.Printf("Length Remaing:%v\n", len(remaing))
	}
	fmt.Println("Done Sieve")
	max := 5000000
	for value := range remaing {
		if value < max {
			max = value
		}
	}

	return max
}

func removeBases(remaining map[int]bool, pos, depth int) {
	period := (depth * 2) - 2
	base := period - pos
	for x := base; x < 5000000; x = x + period {
		delete(remaining, x)
	}

}
