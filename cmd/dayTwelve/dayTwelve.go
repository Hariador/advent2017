package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Mesh graph.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	nodeList := getSourceData(filename)
	count := 0
	var size int
	for len(nodeList) > 0 {
		fmt.Println(nodeList)
		size, nodeList = findNetworkSize(nodeList)
		fmt.Printf("Network size: %v\n", size)
		count++
	}
	fmt.Printf("Number of networks: %v\n", count)

}

func getSourceData(filename string) []node {
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	lineScanner := bufio.NewScanner(reader)
	lineScanner.Split(bufio.ScanLines)
	var nodeList []node
	var tempChildren []string
	for lineScanner.Scan() {
		tempLine := lineScanner.Text()
		wordScanner := bufio.NewScanner(strings.NewReader(tempLine))
		wordScanner.Split(bufio.ScanWords)
		wordScanner.Scan()
		name := wordScanner.Text()
		wordScanner.Scan()
		for wordScanner.Scan() {
			child := wordScanner.Text()
			tempChildren = append(tempChildren, strings.Trim(child, ","))
		}
		newNode := node{name, tempChildren}
		nodeList = append(nodeList, newNode)
		tempChildren = nil
	}

	return nodeList
}

func findNetworkSize(nodes []node) (int, []node) {
	foundNodes := make(map[string]bool)
	workingSet := nodes
	found := true
	var temp []node
	n, workingSet := workingSet[0], workingSet[1:]
	findNode(n, foundNodes)
	for found {
		temp = nil
		found = false
		for _, n := range workingSet {
			if exists(n, foundNodes) {
				findNode(n, foundNodes)
				found = true
			} else {
				temp = append(temp, n)
			}
		}
		workingSet = temp
	}

	return len(foundNodes), workingSet

}

func findNode(n node, foundNodes map[string]bool) {
	foundNodes[n.name] = true
	for _, child := range n.connections {
		foundNodes[child] = true
	}
}

func exists(n node, foundNodes map[string]bool) bool {
	_, state := foundNodes[n.name]
	if state {
		return true
	}
	for _, child := range n.connections {
		_, state := foundNodes[child]
		if state {
			return true
		}
	}

	return false
}

type node struct {
	name        string
	connections []string
}
