package main

import (
	"errors"
	"fmt"
)

type firewall struct {
	walls    map[int]*wall
	max      int
	maxWidth int
	position int
	scan     int
}

func (fw *firewall) Add(w *wall, index int) {
	fw.walls[index] = w
	if index > fw.max {
		fw.max = index
	}
	if w.width > fw.maxWidth {
		fw.maxWidth = w.width
	}
}

func (fw *firewall) step() error {
	fw.position++
	found := false
	wall, ok := fw.walls[fw.position]
	if ok {
		found = wall.Check(1)

	}

	if found {

		return errors.New("was found")

	}
	for _, wall := range fw.walls {
		wall.Move()
	}
	return nil
}

func (fw *firewall) testPass(delay int) bool {
	fw.position = -1 - delay
	for fw.position < fw.max {
		err := fw.step()
		if err != nil {
			return true
		}
	}

	return false
}

func (fw *firewall) Reset() {
	for _, wall := range fw.walls {
		wall.Reset()

	}
}

func (fw *firewall) Print() {
	fmt.Printf("Position: %v\tScan: %v\n", fw.position, fw.scan)
}
