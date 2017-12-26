package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	north = iota
	south
	east
	west
)

const (
	alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	argCount := len(os.Args)
	if argCount < 2 {
		panic("./day19 <filename>")
	}

	filename := os.Args[1]
	var routeMap RouteMap
	routeMap.New(getSourceData(filename))
	routeMap.setCoords(15, 0)
	routeMap.dir = south
	routeMap.steps = 0
	fmt.Printf("Height: %v\tWidth: %v\n", routeMap.height, routeMap.width)
	done := false
	for !done {
		routeMap.PrintStats()
		routeMap.steps++
		done = routeMap.move()
	}
	fmt.Println(routeMap.path)
	fmt.Println(routeMap.steps)
}

func getSourceData(filename string) ([][]string, int, int) {
	f, _ := os.Open(filename)
	var temp [][]string
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	x := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if x == 0 {
			x = len(line)
		}
		runeLine := parseLine(line)
		temp = append(temp, runeLine)
		y++
	}

	return temp, x, y

}

func parseLine(line string) []string {
	var temp []string
	for _, char := range line {
		sChar := string(char)
		temp = append(temp, sChar)
	}

	return temp
}

type RouteMap struct {
	route         [][]string
	x, y          int
	height, width int
	dir           int
	path          string
	steps         int
}

func (r *RouteMap) New(input [][]string, width, height int) {
	r.route = input
	r.height = height
	r.width = width
	r.x = 0
	r.y = 0
}

func (r *RouteMap) Print() {
	for x := 0; x < r.height; x++ {
		fmt.Printf("%v\n", r.route[x])
	}
}

func (r *RouteMap) PrintSub(x, y int) {
	subX := x - 2
	subY := y - 2
	maxX := x + 2
	maxY := y + 2
	if subX < 0 {
		subX = 0
	}
	if subY < 0 {
		subY = 0
	}
	if maxX >= r.width {
		maxX = r.width
	}
	if maxY >= r.height {
		maxY = r.width
	}
	for i := subX; i < maxX; i++ {
		for j := subY; j < maxY; j++ {
			fmt.Print(r.route[i][j])
		}
		fmt.Printf("\n")
	}
}

func (r *RouteMap) PrintStats() {
	fmt.Printf("X: %v\tY:%v\tVALUE: %v\tDIR: %v\n", r.x, r.y, r.get(), r.dir)
}

func (r *RouteMap) setCoords(x, y int) {

	r.x = x
	r.y = y
}

func (r *RouteMap) get() string {
	return r.route[r.y][r.x]
}

func (r *RouteMap) getCharAtCoords(x, y int) string {
	return r.route[y][x]
}

func (r *RouteMap) scan() {
	if r.dir != north && r.y+1 < r.height {
		fmt.Printf("NORTH=X: %v\tY: %v\t H: %v\tW: %v\n", r.x, r.y, r.height, r.width)
		if r.getCharAtCoords(r.x, r.y+1) == "|" || strings.Contains(alpha, r.getCharAtCoords(r.x, r.y+1)) {
			r.dir = south

			return
		}

	}
	if r.dir != south && r.y-1 > 0 {
		fmt.Printf("SOUTH=X: %v\tY: %v\t H: %v\tW: %v\n", r.x, r.y, r.height, r.width)
		if r.getCharAtCoords(r.x, r.y-1) == "|" || strings.Contains(alpha, r.getCharAtCoords(r.x, r.y-1)) {
			r.dir = north
			return
		}
	}
	if r.dir != west && r.x+1 < r.width {
		fmt.Printf("West=X: %v\tY: %v\t H: %v\tW: %v\n", r.x, r.y, r.height, r.width)
		if r.getCharAtCoords(r.x+1, r.y) == "-" || strings.Contains(alpha, r.getCharAtCoords(r.x+1, r.y)) {
			r.dir = east
			return
		}

	}
	if r.dir != east && r.x-1 > 0 {
		fmt.Printf("East=X: %v\tY: %v\t H: %v\tW: %v\n", r.x, r.y, r.height, r.width)
		if r.getCharAtCoords(r.x-1, r.y) == "-" || strings.Contains(alpha, r.getCharAtCoords(r.x-1, r.y)) {
			r.dir = west

			return
		}
	}
}

func (r *RouteMap) check() bool {
	fmt.Println("Checking")
	switch {
	case r.get() == "+":
		{
			r.scan()
		}
	case strings.Contains(alpha, r.get()):
		{
			r.path = r.path + r.get()
		}
	case r.get() == " ":
		{
			return false
		}

	}
	return true
}

func (r *RouteMap) move() bool {
	if !r.check() {
		return true
	}
	switch {
	case r.dir == north:
		{
			r.y--
		}
	case r.dir == south:
		{
			r.y++
		}
	case r.dir == west:
		{
			r.x--
		}
	case r.dir == east:
		{
			r.x++
		}
	}
	return false
}
