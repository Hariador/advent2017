package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argCount := len(os.Args)
	if argCount < 2 {
		panic("In valid syntax or something.  Look, I'm not your mother.")
	}

	stepSize, _ := strconv.Atoi(os.Args[1])
	node := Node{}
	node.value = 0
	node.next = &node
	node.prev = &node
	var value int
	var pos int
	pos = 1
	//lock := Spinlock{stepSize, &node, &node, &node}
	for x := 1; x < 50000000; x++ {
		if x%100000 == 0 {
			fmt.Println(x)
		}
		pos = (pos + stepSize) % x
		pos++
		if pos == 1 {
			value = x

		}
		//fmt.Printf("POS: %v\tVALUE: %v\n", pos, value)

		//lock.Step(x)
		//lock.Print(5)

	}
	//lock.Print(5)
	fmt.Printf("Value: %v\n", value)
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type Spinlock struct {
	stepSize int
	curr     *Node
	tail     *Node
	head     *Node
}

func (s *Spinlock) Step(value int) {
	s.Advance(s.stepSize)
	s.Insert(value)

}

func (s *Spinlock) Insert(value int) {
	var temp *Node
	temp = s.curr.next
	newNode := Node{}
	newNode.value = value
	newNode.prev = s.curr
	newNode.next = temp
	newNode.prev.next = &newNode
	s.curr = &newNode
}

func (s *Spinlock) Advance(amount int) {
	for x := 0; x < amount; x++ {
		s.curr = s.curr.next
	}
}

func (s *Spinlock) Print(amount int) {
	temp := s.head
	for x := 0; x <= amount; x++ {
		if temp == s.curr {
			fmt.Printf("(%v) ", temp.value)
		} else {
			fmt.Printf("%v ", temp.value)
		}
		temp = temp.next
	}
	fmt.Printf("\n")

}
