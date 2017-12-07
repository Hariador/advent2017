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
	sourceData := getSourceData(filename)
	parents, _ := splitList(sourceData)
	root, _ := getRoot(parents)
	fmt.Printf("Root pn name: %v\n", root)
	tower := &ProgramNode{}
	tower.ChildList = make(map[string]ProgramNode)
	tower.build(root, sourceData)
	tower.print(0)
	tower.balance(0)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSourceData(filename string) ProgramList {
	var sourceData ProgramList
	sourceData.List = make(map[string]UnlinkedProgams)
	var tempList []string
	var name string
	var weight int
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tempList = nil
		line := scanner.Text()
		lineReader := bufio.NewReader(strings.NewReader(line))
		wordScanner := bufio.NewScanner(lineReader)
		wordScanner.Split(bufio.ScanWords)

		if wordScanner.Scan() {
			name = wordScanner.Text()
			wordScanner.Scan()
			weight = getWeight(wordScanner.Text())
			//fmt.Printf("Name: %v\tWeight %v\n", name, weight)
		}
		wordScanner.Scan()
		_ = wordScanner.Text()
		for wordScanner.Scan() {

			tempList = append(tempList, strings.Trim(wordScanner.Text(), ","))
		}
		tempNode := UnlinkedProgams{name, weight, tempList}
		sourceData.add(tempNode)
	}

	return sourceData
}

func getWeight(s string) int {
	number := strings.Trim(s, "()")
	result, _ := strconv.Atoi(number)
	return result
}

func splitList(list ProgramList) ([]UnlinkedProgams, []UnlinkedProgams) {
	var childless []UnlinkedProgams
	var parents []UnlinkedProgams
	for name := range list.List {
		item, _ := list.get(name)
		if len(item.Children) > 0 {
			parents = append(parents, item)
		} else {
			childless = append(childless, item)
		}
	}

	return parents, childless
}

func getRoot(list []UnlinkedProgams) (string, bool) {
	var found bool

	for i := 0; i < len(list); i++ {
		found = false
		name := list[i].Name
		for j := 0; j < len(list); j++ {
			if list[j].contains(name) {
				found = true
			}
		}
		if !found {
			return name, true
		}
	}

	return "", false
}

type ProgramNode struct {
	ChildList map[string]ProgramNode
	Weight    int
	Name      string
}

func (pn *ProgramNode) build(name string, list ProgramList) {
	pn.Name = name
	item, _ := list.get(name)

	pn.Weight = item.Weight
	var temp *ProgramNode
	for _, childname := range item.Children {
		temp = new(ProgramNode)
		temp.ChildList = make(map[string]ProgramNode)
		temp.build(childname, list)

		pn.ChildList[childname] = *temp
	}
}

func (pn *ProgramNode) print(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("%v", pn.Name)
	if len(pn.ChildList) > 0 {
		fmt.Print("\n")
	} else {
		fmt.Print("* ")
	}
	for _, child := range pn.ChildList {
		child.print(depth + 1)
	}
	fmt.Print("\n")

}

func (pn *ProgramNode) getWeight() int {
	var aboveWeight int
	for _, p := range pn.ChildList {
		aboveWeight += p.getWeight()
	}

	return pn.Weight + aboveWeight
}

func (pn *ProgramNode) balance(goal int) {
	var weights map[string]int
	unbalanced := false
	weights = make(map[string]int)
	tempWeight := 0
	sum := 0
	if len(pn.ChildList) > 0 {
		for n, p := range pn.ChildList {
			weights[n] = p.getWeight()
			sum += weights[n]
			if tempWeight == 0 {
				tempWeight = weights[n]
			} else {
				if tempWeight != weights[n] {
					unbalanced = true
				}
			}
		}
		fmt.Printf("Weights: %v\t unBalanced: %v\n", weights, unbalanced)
		if unbalanced {
			odd, subGoal, ok := findOddWeight(weights)
			fmt.Printf("Unbalanced Program: %v\tTarget Weight: %v\n", odd, subGoal)
			if ok {
				p := pn.ChildList[odd]
				p.balance(subGoal)
			}
		} else if goal != sum+pn.Weight {
			fmt.Printf("IM GETTING FIXED! Program Name: %v\t Original Weight: %v\tNew Weight: %v\n", pn.Name, pn.Weight, goal-sum)
			pn.Weight = goal - sum
		}
	} else {
		fmt.Printf("IM GETTING FIXED! Program Name: %v\t Original Weight: %v\tNew Weight: %v\n", pn.Name, pn.Weight, goal)
		pn.Weight = goal
	}

}

func findOddWeight(weights map[string]int) (string, int, bool) {
	counts := make(map[int]int)
	var goal int
	if len(weights) == 1 {
		return "", goal, false
	}
	for _, weight := range weights {
		counts[weight]++
		if counts[weight] == 2 {
			goal = weight
		}
	}
	target := 0
	for weight, count := range counts {
		if count == 1 {
			target = weight
		}
	}

	//fmt.Printf("Target:%v\n", target)
	for name, weight := range weights {
		//fmt.Printf("Name: %v\tWeight: %v\n", name, weight)
		if weight == target {
			return name, goal, true
		}
	}
	return "", goal, false
}

type UnlinkedProgams struct {
	Name     string
	Weight   int
	Children []string
}

type ProgramList struct {
	List map[string]UnlinkedProgams
}

func (pl *ProgramList) add(item UnlinkedProgams) {
	name := item.Name
	pl.List[name] = item
}

func (pl *ProgramList) get(name string) (UnlinkedProgams, bool) {
	item, found := pl.List[name]
	return item, found
}

func (up *UnlinkedProgams) contains(name string) bool {
	for i := 0; i < len(up.Children); i++ {
		if name == up.Children[i] {
			return true
		}
	}

	return false
}
