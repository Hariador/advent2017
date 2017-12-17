package main

import (
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
	argCount := len(os.Args)
	if argCount < 2 {
		panic("./day15 <seed1> <seed2>")
	}
	seed1, _ := strconv.Atoi(os.Args[1])
	mult1, _ := strconv.Atoi(os.Args[2])
	seed2, _ := strconv.Atoi(os.Args[3])
	mult2, _ := strconv.Atoi(os.Args[4])
	stop, _ := strconv.Atoi(os.Args[5])
	genOne := generator{int64(seed1), int64(seed1), 0, 16807, int64(mult1)}
	genTwo := generator{int64(seed2), int64(seed2), 9, 48271, int64(mult2)}
	match := 0
	for x := 0; x < stop; x++ {
		genOne.GitGud()

		genTwo.GitGud()
		//	genOne.Print()
		//	genTwo.Print()

		if genOne.LeastSixteen() == genTwo.LeastSixteen() {
			match++
			//genOne.Print()
			//genTwo.Print()

		}
	}

	fmt.Printf("There were %v matches generated\n", match)

}

type generator struct {
	seed       int64
	curr       int64
	generation int
	factor     int64
	mult       int64
}

func (g *generator) Tick() {
	temp := g.curr * g.factor

	rem := temp % 2147483647

	g.curr = rem
	g.generation++
}

func (g *generator) GitGud() {
	g.Tick()
	for g.curr%g.mult != 0 {
		g.Tick()

	}

}

func (g *generator) LeastSixteen() int64 {
	return g.curr & 65535
}

func (g *generator) Print() {
	fmt.Printf("INT: %v\tBIN: %032b\n", g.curr, g.curr)

}
