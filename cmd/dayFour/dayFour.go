package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/willf/bloom"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day Four. pass phrase validation.")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	count := validatePassList(filename)
	fmt.Printf("There are %v valid phrases\n", count)

}

func validatePassList(filename string) int {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	var count int
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineReader := bufio.NewReader(strings.NewReader(line))
		wordScanner := bufio.NewScanner(lineReader)
		wordScanner.Split(bufio.ScanWords)
		if validatePhrase(wordScanner) {
			count++
		}
	}
	return count
}

func validatePhrase(wordScanner *bufio.Scanner) bool {
	filter := bloom.New(1024, 25)
	for wordScanner.Scan() {
		password := wordScanner.Text()
		password = sortString(password)
		if filter.TestString(password) {
			return false
		}

		filter.AddString(password)
	}

	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
