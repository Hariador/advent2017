package main

import "fmt"

import "strconv"
import "os"
import "bufio"
import "strings"

func main() {
	pic := Picture{}
	pic.New(2188)
	interations, _ := strconv.Atoi(os.Args[1])
	var rules []Rule
	filename := os.Args[2]
	rules = getSourceData(filename)
	pic.Init()
	pic.rules = rules
	for x := 0; x < interations; x++ {
		pic.Split()
		fmt.Println("Tick")
		//pic.Print()
	}

	//pic.Print()
	fmt.Println(pic.Count())

}

func getSourceData(filename string) []Rule {
	var temp []Rule
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		ruleParts := strings.Split(parts[0], "/")
		actionParts := strings.Split(parts[1], "/")
		tempRule := Rule{len(ruleParts), ruleParts, actionParts}
		temp = append(temp, tempRule)
	}

	return temp
}

type Rule struct {
	size     int
	patterns []string
	action   []string
}

func (r *Rule) Print() {
	for x := 0; x < r.size; x++ {
		fmt.Println(r.patterns[x])
	}
}
func (r *Rule) PrintOutput() {
	for x := 0; x < r.size+1; x++ {
		fmt.Println(r.action[x])
	}
}

type Picture struct {
	size  int
	frame [][]string
	rules []Rule
}

func (p *Picture) Print() {
	for y := 0; y < p.size; y++ {
		for x := 0; x < p.size; x++ {
			fmt.Print(p.frame[y][x])
		}

		fmt.Printf("\n")

	}

}

func (p *Picture) setLine(line string, x, y int) {
	length := len(line)
	for i := 0; i < length; i++ {
		p.frame[y][x+i] = string(line[i])
	}
}

func (p *Picture) getLine(x, y, length int) string {
	temp := ""
	for i := 0; i < length; i++ {
		temp = temp + p.frame[y][x+i]
	}

	return temp
}

func (p *Picture) setSquare(square Picture, x, y int) {
	for i := 0; i < square.size; i++ {
		for j := 0; j < square.size; j++ {
			p.frame[i+y][j+x] = square.frame[i][j]
		}
	}
}

func (p *Picture) getSquare(x, y, size int) Picture {
	temp := Picture{}
	temp.New(size)
	temp.size = size
	for j := 0; j < size; j++ {
		temp.setLine(p.getLine(x, y+j, size), 0, j)
	}
	return temp
}
func (p *Picture) Init() {
	p.setLine(".#.", 0, 0)
	p.setLine("..#", 0, 1)
	p.setLine("###", 0, 2)
	p.size = 3
}

func (p *Picture) New(size int) {
	p.frame = make([][]string, size)
	for x := 0; x < size; x++ {
		p.frame[x] = make([]string, size)
	}
}

func (p *Picture) Split() {
	increase := 0
	size := 0
	temp := Picture{}
	oldSize := p.size
	if p.size%2 == 0 {
		increase = p.size / 2
		p.size = p.size + increase
		size = 2

	} else {
		increase = p.size / 3
		p.size = p.size + increase
		size = 3

	}
	temp.New(p.size + increase)
	yOffest := 0
	xOffset := 0
	count := 0
	temp.size = p.size
	for i := 0; i < oldSize; i += size {
		xOffset = 0
		count = 0
		for j := 0; j < oldSize; j += size {
			count++
			newSquare := p.getSquare(j, i, size)
			newSquare.rules = p.rules
			newSquare.Transform()

			temp.setSquare(newSquare, j+xOffset, i+yOffest)
			xOffset++

		}
		yOffest++
	}

	p.setSquare(temp, 0, 0)

}

func (p *Picture) Count() int {
	count := 0
	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			if p.frame[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func (p *Picture) HFlip() {
	temp := Picture{}
	temp.New(p.size)
	temp.size = p.size
	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			temp.frame[i][j] = p.frame[p.size-i-1][j]
		}
	}
	p.setSquare(temp, 0, 0)
}

func (p *Picture) Vflip() {
	temp := Picture{}
	temp.New(p.size)
	temp.size = p.size
	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			temp.frame[i][j] = p.frame[i][p.size-j-1]
		}
	}
	p.setSquare(temp, 0, 0)

}

func (p *Picture) ClockWise() {
	temp := Picture{}
	temp.New(p.size)
	temp.size = p.size
	for i := 0; i < p.size; i++ {
		for j := 0; j < p.size; j++ {
			temp.frame[i][j] = p.frame[p.size-j-1][i]
		}
	}
	p.setSquare(temp, 0, 0)
}

func (p *Picture) Transform() {
	var match bool
	temp := p
	for _, rule := range p.rules {
		match = true
		if temp.size == rule.size {

			match = compare(rule, temp)
			if !match {
				temp.ClockWise()
				match = compare(rule, temp)

			}
			if !match {

				temp.ClockWise()

				match = compare(rule, temp)
			}
			if !match {
				temp.ClockWise()

				match = compare(rule, temp)

			}
			if !match {
				temp = p
				temp.Vflip()
				match = compare(rule, temp)

			}
			if !match {
				temp.ClockWise()
				match = compare(rule, temp)

			}
			if !match {

				temp.ClockWise()

				match = compare(rule, temp)
			}
			if !match {
				temp.ClockWise()

				match = compare(rule, temp)

			}

			if !match {
				temp = p
				temp.HFlip()

				match = compare(rule, temp)
			}
			if !match {
				temp.ClockWise()
				match = compare(rule, temp)

			}
			if !match {

				temp.ClockWise()

				match = compare(rule, temp)
			}
			if !match {
				temp.ClockWise()

				match = compare(rule, temp)

			}

			if match {

				p.size++
				p.New(p.size)
				for x := 0; x < rule.size+1; x++ {

					p.setLine(rule.action[x], 0, x)

				}

				return
			}
		}
	}
	panic("OH SHIT")
}

func compare(rule Rule, p *Picture) bool {

	for x := 0; x < rule.size; x++ {
		test := rule.patterns[x]
		sub := p.getLine(0, x, p.size)
		//	fmt.Printf("COMPARING: %v\t%v\t%v\n", test, sub, strings.Compare(test, sub))
		if strings.Compare(test, sub) != 0 {
			return false
		}
	}
	return true
}
