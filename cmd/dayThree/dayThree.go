package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day two. Spreadsheet check sum.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	spot, _ := strconv.Atoi(filename)
	distance := rectilinearDistance(spot)
	_ = renderUntil(spot)
	fmt.Printf("Retrival Distance: %v\n", distance)
}

func rectilinearDistance(spot int) float64 {
	x, y := getCoords(spot)
	dist := math.Abs(float64(x)) + math.Abs(float64(y))

	return dist
}

func getCoords(spot int) (int, int) {
	fmt.Println(spot)
	var x int
	var y int
	root := math.Sqrt(float64(spot))
	power := int(math.Ceil(root))

	if power%2 == 0 {
		power++
	}
	lastPower := power - 2
	minCellValue := (lastPower * lastPower) + 1
	maxCellValue := power * power
	diff := maxCellValue - spot
	mid := ((maxCellValue - minCellValue) + 1) / 2
	right := mid + power - 1
	switch {
	case diff <= power:
		{
			x = power - diff - 1
			y = power - 1
		}
	case diff <= mid:
		{
			x = 0
			y = mid - diff
		}
	case diff <= right:
		{
			x = diff - mid
			y = 0
		}

	case diff > right:
		{
			x = power - 1
			y = diff - right
		}

	}
	shift := (power - 1) / 2
	x = x - shift
	y = y - shift
	fmt.Printf("Address: %v\tPower:%v\tMax: %v\tMin: %v\tDiff: %v\tMid:%v\tRight: %v\tX: %v\tY: %v\n", spot, power, maxCellValue, minCellValue, diff, mid, right, x, y)
	return x, y
}

func renderUntil(spot int) int {
	fmt.Printf("Find first value larger than %v\n", spot)
	var addressField = [11][11]int{}
	shift := 5
	x := shift
	y := shift
	var value int
	right := true
	up := false
	left := false
	down := false
	edge := 1
	addressField[x][y] = 1
	value = addressField[x][y]
	fmt.Println(addressField)
	for value < spot {
		fmt.Printf("Value: %v\tSpot: %v\tX: %v\tY: %v\tEdge: %v\n", value, spot, x, y, edge)
		switch {
		case right:
			{
				x++
				value = addAround(addressField, x, y)
				addressField[x][y] = value
				if x-shift == edge {
					turn(&right, &up, &left, &down)
				}

			}
		case up:
			{
				y++
				value = addAround(addressField, x, y)
				addressField[x][y] = value
				if y-shift == edge {
					turn(&right, &up, &left, &down)
					edge = edge * -1
				}
			}
		case left:
			{
				x--
				value = addAround(addressField, x, y)
				addressField[x][y] = value
				if x-shift == edge {
					turn(&right, &up, &left, &down)
				}
			}
		case down:
			{
				y--
				value = addAround(addressField, x, y)
				addressField[x][y] = value
				if y-shift == edge {
					turn(&right, &up, &left, &down)
					edge = (edge * -1) + 1
				}
			}
		}

	}
	draw(addressField)
	return shift
}

func addAround(addressField [11][11]int, x int, y int) int {
	value := 0
	fmt.Printf("X: %v\tY: %v\n", x, y)
	value = value + addressField[x-1][y-1]
	value = value + addressField[x-1][y]
	value = value + addressField[x-1][y+1]
	value = value + addressField[x][y+1]
	value = value + addressField[x+1][y+1]
	value = value + addressField[x+1][y]
	value = value + addressField[x+1][y-1]
	value = value + addressField[x][y-1]

	return value
}

func turn(right *bool, up *bool, left *bool, down *bool) {
	fmt.Println("TURNING")
	switch {
	case *right:
		{
			*right = false
			*up = true
		}
	case *up:
		{
			*up = false
			*left = true
		}
	case *left:
		{
			*left = false
			*down = true
		}
	case *down:
		{
			*down = false
			*right = true
		}
	}

}

func draw(addressField [11][11]int) {
	for y := 0; y < 11; y++ {
		for x := 0; x < 11; x++ {
			fmt.Printf("%v\t", addressField[x][y])
		}
		fmt.Printf("\n")
	}
}
