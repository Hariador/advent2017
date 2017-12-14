package main

import "fmt"

type wall struct {
	curr     int
	width    int
	outgoing bool
}

func (w *wall) Move() {
	if w.outgoing {
		if w.curr == w.width {
			w.outgoing = false
			w.curr--
		} else {
			w.curr++
		}
	} else {
		if w.curr == 1 {
			w.outgoing = true
			w.curr++
		} else {
			w.curr--
		}
	}
}

func (w *wall) New(width int) wall {
	temp := wall{}
	temp.width = width
	temp.curr = 1
	temp.outgoing = true
	return temp
}

func (w *wall) Check(pos int) bool {
	if w.curr == pos {
		return true
	}

	return false
}

func (w *wall) Print() {
	fmt.Printf("Curr: %v\tWidth: %v\tOut: %v\n", w.curr, w.width, w.outgoing)
}

func (w *wall) Reset() {
	w.curr = 0
	w.outgoing = true
}
