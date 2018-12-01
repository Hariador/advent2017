package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	north = iota
	east
	south
	west
)

const (
	clean = iota
	weakened
	infected
	flagged
)

const (
	white = "\033[1;37m"
	green = "\033[1;32m"
	red   = "\033[1;31m"
	on    = "\033[7;31m"
	off   = "\033[0;31m"
	yello = "\033[1;33m"
	teal  = "\033[1;34m"
)

func main() {
	filename := os.Args[1]
	iterations, _ := strconv.Atoi(os.Args[2])
	makePretty := false
	if len(os.Args) == 4 {
		prettyFlag := os.Args[3]
		if prettyFlag == "-P" || prettyFlag == "-p" {
			makePretty = true
		}
	}

	memory := getSourceData(filename)
	memory.print()
	mS := time.Millisecond
	for x := 0; x < iterations; x++ {
		memory.burst()
		if makePretty {
			time.Sleep(mS * 75)
			memory.print()
		}
	}
	fmt.Printf("Infects: %v\n", memory.infects)

}

func getSourceData(filename string) Grid {
	temp := Grid{}
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	firstLine := scanner.Text()
	if len(firstLine)%2 == 0 {
		panic("CAN'T HAVE A MIDDLE IF IT's EvEn")
	}
	temp.size = len(firstLine)
	temp.offset = temp.size / 2
	temp.vX = temp.offset
	temp.vY = temp.offset
	temp.New(temp.size)
	lineScanner := bufio.NewScanner(strings.NewReader(firstLine))
	lineScanner.Split(bufio.ScanBytes)
	x := 0
	y := 0
	for lineScanner.Scan() {
		if lineScanner.Text() == "." {
			temp.set(x, y, clean)
		} else {
			temp.set(x, y, infected)
		}
		x++
	}
	for scanner.Scan() {
		x = 0
		y++
		firstLine = scanner.Text()
		lineScanner := bufio.NewScanner(strings.NewReader(firstLine))
		lineScanner.Split(bufio.ScanBytes)
		for lineScanner.Scan() {
			if lineScanner.Text() == "." {
				temp.set(x, y, clean)
			} else {
				temp.set(x, y, infected)
			}
			x++
		}
	}

	return temp
}

type Grid struct {
	size     int
	offset   int
	cells    [][]int
	virusDir int
	vX, vY   int
	infects  int
}

func (g *Grid) New(size int) {
	g.cells = make([][]int, size)
	for x := 0; x < size; x++ {
		g.cells[x] = make([]int, size)
	}
}

func (g *Grid) grow() {

	temp := make([][]int, g.size+4)
	for y := 0; y < g.size+4; y++ {
		temp[y] = make([]int, g.size+4)
		for x := 0; x < g.size+4; x++ {
			temp[y][x] = clean
		}
	}
	swap := g.cells
	g.cells = temp
	g.mapSubGrid(2, 2, swap, g.size)
	g.size += 4
	g.vX += 2
	g.vY += 2

}

func (g *Grid) burst() {
	if g.vX <= 2 || g.vX > g.size-2 || g.vY <= 2 || g.vY > g.size-2 {
		g.grow()
	}

	node := g.get(g.vX, g.vY)
	if node == weakened {
		g.infects++
	}
	switch {
	case node == clean:
		{
			g.virusDir--
			if g.virusDir < 0 {
				g.virusDir = 3
			}
		}
	case node == weakened:
		{
		}
	case node == infected:
		{
			g.virusDir = (g.virusDir + 1) % 4
		}
	case node == flagged:
		{
			g.virusDir += 2
			if g.virusDir > 3 {
				g.virusDir -= 4
			}
		}
	}
	node = (node + 1) % 4
	g.set(g.vX, g.vY, node)
	switch {
	case g.virusDir == north:
		{
			g.vY--
		}
	case g.virusDir == south:
		{
			g.vY++
		}
	case g.virusDir == east:
		{
			g.vX++
		}
	case g.virusDir == west:
		{
			g.vX--
		}
	}
}

func (g *Grid) mapSubGrid(x, y int, subGrid [][]int, size int) {
	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			g.set(j, i, subGrid[i-y][j-x])
		}
	}
}

func (g *Grid) get(x, y int) int {
	return g.cells[y][x]
}

func (g *Grid) set(x, y int, value int) {

	g.cells[y][x] = value
}

func (g *Grid) print() {
	fmt.Printf("\033[2J")
	fmt.Printf("\033[1;1H")
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			node := g.cells[y][x]
			switch {
			case node == clean:
				{
					fmt.Print(green)
					fmt.Print(".")
				}
			case node == infected:
				{
					fmt.Print(red)
					fmt.Print("#")
				}
			case node == weakened:
				{
					fmt.Printf("%vW", yello)
				}
			case node == flagged:
				{
					fmt.Printf("%vF", teal)
				}

			}

		}
		fmt.Println()
	}
	col := g.vX + 1
	row := g.vY + 1
	fmt.Printf("\033[%v;%vH", row, col)
	fmt.Printf("%v☺%v", on, off)
	fmt.Printf("\033[%v;%vH", g.size+1, g.size+1)
	fmt.Print(white)
	fmt.Println()
}
