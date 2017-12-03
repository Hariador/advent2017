package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Advent 2016 Day one A \nReverse Captcha")

	argCount := len(os.Args)
	if argCount < 2 {
		panic("Need to have a filename as a second argument")
	}

	filename := os.Args[1]
	fmt.Printf("Using %s as source data\n", filename)
	answerOne, answerTwo, err := arraySolve(filename)
	check(err)
	fmt.Println("*************************************")
	fmt.Println(answerOne)
	fmt.Println(answerTwo)

}

func arraySolve(filename string) (int, int, error) {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	rawDataSet, readBytes := getLongNumberList(f)
	var sumOne int
	var sumTwo int
	var i int
	for i = 0; i < readBytes; i++ {
		//Convert from asci to int

		if rawDataSet[i] == rawDataSet[(i+1)%readBytes] {
			sumOne = sumOne + int(rawDataSet[i]) - 48
		}
		if rawDataSet[i] == rawDataSet[(i+(readBytes/2))%readBytes] {
			sumTwo = sumTwo + int(rawDataSet[i]) - 48
		}

	}
	// if rawDataSet[i] == rawDataSet[0] {
	// 	sumOne = sumOne + int(rawDataSet[0]) - 48
	// }

	return sumOne, sumTwo, nil
}

func getLongNumberList(f *os.File) ([]byte, int) {
	fileInfo, err := f.Stat()
	check(err)
	dataSetLength := fileInfo.Size()
	rawDataSet := make([]byte, dataSetLength-1)
	readBytes, err := f.Read(rawDataSet)

	return rawDataSet, readBytes
}
