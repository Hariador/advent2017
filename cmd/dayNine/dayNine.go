package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day Nine, stream calculator.")
	argCount := len(os.Args)
	if argCount < 1 {
		panic("Need to have a filename to open as arg1")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	check(err)
	reader := bufio.NewReader(f)
	_, _, _ = reader.ReadRune()
	score, length, err := group(reader, 0)
	check(err)
	fmt.Printf("Total score for the stream is %v and we removed %v characters of garbage.", score, length)
}

func group(reader *bufio.Reader, score int) (int, int, error) {
	groupScore := score + 1
	totalLength := 0
	end := false
	for !end {
		currRune, _, err := reader.ReadRune()
		now := string(currRune)

		if err != nil {
			panic("Never found the matching closing brace")
		}
		switch {
		case now == "{":
			{
				subScore, length, err := group(reader, score+1)
				groupScore += subScore
				totalLength += length
				check(err)
			}
		case now == "}":
			{
				end = true
			}
		case now == "<":
			{
				length, err := garbage(reader)
				check(err)
				fmt.Printf("Consumed a %v character garbage string.\n", length)
				totalLength += length
			}
		}

	}
	fmt.Printf("Subscore: %v\n", groupScore)
	return groupScore, totalLength, nil
}

func garbage(reader *bufio.Reader) (int, error) {
	length := 0
	end := false
	for !end {
		currRune, _, err := reader.ReadRune()
		check(err)
		now := string(currRune)
		switch {
		case now == "!":
			{
				_, _, err = reader.ReadRune()
				check(err)
			}
		case now == ">":
			{
				end = true
			}
		default:
			{
				length++
			}
		}
	}
	return length, nil
}
