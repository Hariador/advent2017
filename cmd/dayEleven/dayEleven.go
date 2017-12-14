package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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
	inputs := getSourceData(filename)
	hex := &hexGrid{0, 0, 0}
	total := 0
	for _, input := range inputs {
		hex.move(input)
		if total < hex.Dist() {
			total = hex.Dist()
		}

	}
	fmt.Printf("Child program is %v away and was maximally %v steps away\n", hex.Dist(), total)

}

func getSourceData(filename string) []string {

	bytes, _ := ioutil.ReadFile(filename)
	line := string(bytes)

	return strings.Split(line, ",")

}

func (h *hexGrid) move(direction string) {
	switch {
	case direction == "n":
		{
			h.N()
		}
	case direction == "s":
		{
			h.S()
		}
	case direction == "ne":
		{
			h.NE()
		}
	case direction == "se":
		{
			h.SE()
		}
	case direction == "nw":
		{
			h.NW()
		}
	case direction == "sw":
		{
			h.SW()
		}
	}

}

type hexGrid struct {
	x int
	y int
	z int
}

func (h *hexGrid) N() {
	h.y++
	h.z--
}

func (h *hexGrid) S() {
	h.y--
	h.z++
}

func (h *hexGrid) NE() {
	h.x++
	h.z--
}

func (h *hexGrid) SW() {
	h.x--
	h.z++
}

func (h *hexGrid) SE() {
	h.x++
	h.y--
}

func (h *hexGrid) NW() {
	h.x--
	h.y++
}

func (h *hexGrid) Dist() int {
	absX := math.Abs(float64(h.x))
	absY := math.Abs(float64(h.y))
	absZ := math.Abs(float64(h.z))
	dist := (absX + absY + absZ) / 2

	return int(dist)
}

func (h *hexGrid) print() {
	fmt.Printf("(%v, %v,%v)\n", h.x, h.y, h.z)
}
