package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argLength := len(os.Args)
	if argLength < 2 {
		panic("./day18 <filename>")
	}

	filename := os.Args[1]
	ins := getSourceData(filename)
	programOne := runSpace{}
	programOne.init(ins)
	programOne.name = "One"
	programZero := runSpace{}
	programZero.init(ins)
	programZero.name = "Zero"

	chanA := make(chan int, 100000)
	chanB := make(chan int, 100000)
	tickcount := 0
	programOne.registers["p"] = 1
	programZero.initChannels(chanA, chanB)
	programOne.initChannels(chanB, chanA)
	for !(programOne.waiting && programZero.waiting) {
		programZero.tick()
		programOne.tick()
		tickcount++
		if tickcount > 2800 {
			fmt.Println(programZero.registers)
			panic("Loop")
		}
		//fmt.Printf("One sent %v\tZero: %v\tOne: %v\n", programOne.sent, len(programZero.sendBuf), len(programZero.rcvBuf))
	}
	fmt.Printf("Found frequence: %v\n", tickcount)

}

func getSourceData(filename string) []instruction {
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var instructions []instruction
	for scanner.Scan() {
		line := scanner.Text()
		nibbles := strings.Split(line, " ")
		ins := instruction{}
		ins.op = nibbles[0]
		ins.reg = nibbles[1]
		if len(nibbles) > 2 {
			ins.value = nibbles[2]
		}
		instructions = append(instructions, ins)
	}

	return instructions
}

type instruction struct {
	op    string
	reg   string
	value string
}

type runSpace struct {
	registers    map[string]int
	freq         int
	instructions []instruction
	ptr          int
	rcvBuf       <-chan int
	sendBuf      chan<- int
	waiting      bool
	sent         int
	name         string
	max          int
}

func (rS *runSpace) init(ins []instruction) {
	rS.ptr = 0
	rS.freq = 0
	rS.instructions = ins
	rS.waiting = false
	rS.sent = 0
	rS.max = len(rS.instructions)
	rS.registers = make(map[string]int)
	for x := 0; x < 26; x++ {
		rS.registers[string(x+97)] = 0
	}

}

func (rS *runSpace) initChannels(send chan<- int, rcv <-chan int) {
	rS.sendBuf = send
	rS.rcvBuf = rcv
}

func (rS *runSpace) tick() {

	ins := rS.instructions[rS.ptr]
	// if rS.name == "One" {
	// 	fmt.Printf("Program: %v\tINS: %v\tPTR: %v\n", rS.name, ins, rS.ptr)
	// 	fmt.Println(rS.registers)
	// }
	switch {
	case ins.op == "snd":
		rS.send(ins)
	case ins.op == "set":
		rS.set(ins)
	case ins.op == "add":
		rS.add(ins)
	case ins.op == "mul":
		rS.multiply(ins)
	case ins.op == "mod":
		rS.mod(ins)
	case ins.op == "rcv":
		{
			rS.receive(ins)
		}
	case ins.op == "jgz":
		rS.jump(ins)

	}
	rS.ptr++

}

func (rS *runSpace) send(in instruction) {
	rS.sent++
	value, err := strconv.Atoi(in.value)
	if err == nil {
		rS.sendBuf <- value
	} else {
		rS.sendBuf <- rS.registers[in.value]
	}
}

func (rS *runSpace) set(in instruction) {
	value, err := strconv.Atoi(in.value)
	if err == nil {
		rS.registers[in.reg] = value
	} else {
		rS.registers[in.reg] = rS.registers[in.value]
	}

}

func (rS *runSpace) add(in instruction) {
	if in.reg == "i" && rS.name == "Zero" {
		fmt.Println(rS.registers[in.reg])
	}
	value, err := strconv.Atoi(in.value)
	if err == nil {
		rS.registers[in.reg] += value
	} else {

	}
	rS.registers[in.reg] += rS.registers[in.value]
}

func (rS *runSpace) multiply(in instruction) {
	value, _ := strconv.Atoi(in.value)
	value, err := strconv.Atoi(in.value)
	if err == nil {
		rS.registers[in.reg] = rS.registers[in.reg] * value
	} else {
		rS.registers[in.reg] = rS.registers[in.reg] * rS.registers[in.value]
	}
}

func (rS *runSpace) mod(in instruction) {

	value, err := strconv.Atoi(in.value)
	if err == nil {
		fmt.Printf("MOD:%v\t %v:%v\n", rS.registers[in.reg]%value, rS.registers[in.reg], value)
		rS.registers[in.reg] = rS.registers[in.reg] % value
	} else {
		rS.registers[in.reg] = rS.registers[in.reg] % rS.registers[in.value]
	}

}

func (rS *runSpace) receive(in instruction) {
	select {
	case value, err := <-rS.rcvBuf:
		{
			fmt.Printf("Name: %v\t%v\n", rS.name, err)
			rS.registers[in.reg] = value
			rS.waiting = false
		}
	default:
		{
			fmt.Printf("Waiting: %v\n", rS.name)
			rS.ptr--
			rS.waiting = true
		}
	}

}

func (rS *runSpace) jump(in instruction) {
	reg, err := strconv.Atoi(string(in.reg))
	jump := false
	if err == nil {
		if reg > 0 {
			jump = true
		}
	} else {
		if rS.registers[in.reg] > 0 {
			jump = true
		}
	}
	if jump {
		value, err := strconv.Atoi(in.value)
		if err == nil {
			rS.ptr += value
		} else {
			rS.ptr += rS.registers[in.value]
		}
		rS.registers[in.reg] += rS.registers[in.value]
		rS.ptr--
	}
}
