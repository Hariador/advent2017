package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day two. Spreadsheet check sum.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	instructions, registers := getSourceData(filename)
	fmt.Println(instructions)
	fmt.Println(registers)
	maxEVAR := 0
	for _, ins := range instructions {

		ins.Execute(registers)
		curr := findLargest(registers)
		if curr > maxEVAR {
			maxEVAR = curr
		}
	}
	max := findLargest(registers)
	fmt.Printf("MAX: %v\tMAX EVAR: %v\n", max, maxEVAR)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSourceData(filename string) ([]Instruction, map[string]int) {
	var instuctionSet []Instruction
	registers := make(map[string]int)

	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		ins := processLine(line)
		_, ok := registers[ins.reg]
		if !ok {
			registers[ins.reg] = 0
		}
		instuctionSet = append(instuctionSet, ins)
	}

	return instuctionSet, registers
}

func processLine(input string) Instruction {

	lineReader := bufio.NewReader(strings.NewReader(input))
	wordScanner := bufio.NewScanner(lineReader)
	wordScanner.Split(bufio.ScanWords)
	wordScanner.Scan()
	register := wordScanner.Text()
	wordScanner.Scan()
	op := wordScanner.Text()
	wordScanner.Scan()
	value, _ := strconv.Atoi(wordScanner.Text())
	wordScanner.Scan() // Eat the If
	wordScanner.Scan()
	target := wordScanner.Text()
	wordScanner.Scan()
	condition := wordScanner.Text()
	wordScanner.Scan()
	conValue, _ := strconv.Atoi(wordScanner.Text())

	ins := Instruction{register, op, value, target, condition, conValue}
	return ins
}

func findLargest(registers map[string]int) int {
	max := 0
	for _, value := range registers {
		if value > max {
			max = value
		}
	}

	return max
}

type Instruction struct {
	reg       string
	op        string
	value     int
	tar       string
	condition string
	conValue  int
}

func (ins *Instruction) Execute(registers map[string]int) {
	ins.print()
	if runCondition(ins.condition, registers[ins.tar], ins.conValue) {

		do(ins.op, ins.value, ins.reg, registers)
	}
}

func (ins *Instruction) print() {
	fmt.Printf("REG:%v\tOP:%v\tVL:%v\tTAR:%v\tCON:%v\tCONV:%v\n", ins.reg, ins.op, ins.value, ins.tar, ins.condition, ins.conValue)
}

func runCondition(con string, A int, B int) bool {
	switch {
	case con == ">":
		{
			if A > B {
				return true
			}
		}
	case con == "<":
		{
			if A < B {
				return true
			}
		}
	case con == "<=":
		{
			if A <= B {
				return true
			}
		}
	case con == ">=":
		{
			if A >= B {
				return true
			}
		}
	case con == "==":
		{
			if A == B {
				return true
			}
		}
	case con == "!=":
		{
			if A != B {
				return true
			}
		}

	}

	return false
}

func do(op string, value int, reg string, registers map[string]int) {
	switch {
	case op == "dec":
		{
			registers[reg] = registers[reg] - value
		}
	case op == "inc":
		{
			registers[reg] = registers[reg] + value
		}
	}
}
