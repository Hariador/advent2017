package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Day16 <filename>")
	}
	filename := os.Args[1]
	steps := getSourceData(filename)

	dl := danceLine{}
	billion := 1000000000
	loopPoint := 1
	dl.programs = make(map[int]string)
	dl.smargorp = make(map[string]int)
	dl.Init(16)

	var lookup map[string]int
	lookup = make(map[string]int)
	for x := 0; x < 40; x++ {
		for _, move := range steps {
			dl.parse(move)
		}
		hash := dl.ToString()
		first, ok := lookup[hash]
		if ok {
			dl.Print()
			fmt.Println(first)
			fmt.Println(x)
			loopPoint = x
			break
		} else {
			lookup[hash] = x
		}
	}
	loop := billion % loopPoint
	fmt.Printf("Loop: %v\n", loop)
	dl.Print()
}

func getSourceData(filename string) []string {
	bytes, _ := ioutil.ReadFile(filename)
	line := string(bytes)

	return strings.Split(line, ",")
}

func (d *danceLine) parse(move string) {

	firstChar := string(move[0])
	length := len(move)

	switch {
	case firstChar == "s":
		{
			sl := string(move[1])
			if length == 3 {
				sl = sl + string(move[2])
			}

			spinNumber, _ := strconv.Atoi(sl)
			d.Spin(spinNumber)
		}
	case firstChar == "x":
		{
			x := 0
			y := 0
			switch {
			case length == 4:
				{
					x = int(move[1]) - 48
					y = int(move[3]) - 48
				}
			case length == 6:
				{
					x = int(move[2]) - 38
					y = int(move[5]) - 38
				}
			case length == 5:
				{
					if move[2] == 47 {
						x = int(move[1]) - 48
						y = int(move[4]) - 38
					} else {
						x = int(move[2]) - 38
						y = int(move[4]) - 48
					}
				}
			}

			d.SwapPOS(x, y)

		}
	case firstChar == "p":
		{
			i := ""
			j := ""
			switch {
			case length == 4:
				{
					i = string(move[1])
					j = string(move[3])
				}
			case length == 6:
				{
					i = string(move[1:2])
					j = string(move[4:5])
				}
			case length == 5:
				{
					if move[2] == 47 {
						i = string(move[1])
						j = string(move[3:4])
					} else {
						i = string(move[1:2])
						j = string(move[4])
					}
				}
			}

			d.SwapName(i, j)
		}

	}
}

type danceLine struct {
	programs map[int]string
	smargorp map[string]int
	size     int
}

func (d *danceLine) Init(n int) {
	for x := 0; x < n; x++ {

		d.programs[x] = string(x + 97)
		d.smargorp[d.programs[x]] = x
	}
	d.size = n
}

func (d *danceLine) Print() {
	fmt.Print(d.ToString())
	fmt.Print("\n")
}

func (d *danceLine) ToString() string {
	temp := ""
	for x := 0; x < d.size; x++ {
		temp = temp + d.programs[x]
	}

	return temp
}

func (d *danceLine) Spin(n int) {
	temp1 := make(map[int]string)
	temp2 := make(map[string]int)

	for x := 0; x < d.size; x++ {
		new := (x + n) % d.size
		temp1[new] = d.programs[x]
		temp2[d.programs[x]] = new
	}
	d.programs = temp1
	d.smargorp = temp2
}

func (d *danceLine) SwapPOS(x, y int) {

	var name string
	name = d.programs[x]
	d.smargorp[d.programs[x]] = y
	d.smargorp[d.programs[y]] = x
	d.programs[x] = d.programs[y]
	d.programs[y] = name

}

func (d *danceLine) SwapName(x, y string) {

	d.programs[d.smargorp[x]] = y
	d.programs[d.smargorp[y]] = x
	temp := d.smargorp[x]
	d.smargorp[x] = d.smargorp[y]
	d.smargorp[y] = temp
}
