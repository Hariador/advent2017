package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day two. Spreadsheet check sum.")
	argCount := len(os.Args)
	if argCount < 3 {
		panic("dayTen <array length> <filename>")
	}

	size := os.Args[1]
	fmt.Printf("Size of %v\n", size)
	filename := os.Args[2]
	fmt.Printf("Using %s as source data\n", filename)
	s, _ := strconv.Atoi(size)
	array := initArray(s)
	inputs := getInputs(filename)

	hash := encrypt(array, inputs)
	fmt.Println(string(hash))

}

func initArray(size int) []byte {
	array := make([]byte, size)
	for k := range array {
		array[k] = byte(k)
	}
	return array

}

func getInputs(filename string) []byte {
	f, err := os.Open(filename)
	var inputs []byte

	check(err)
	reader := bufio.NewReader(f)
	line, _, _ := reader.ReadLine()
	inputs = append(line, 17, 31, 73, 47, 23)
	return inputs
}

func reverse(array []byte) []byte {
	var temp []byte
	for i := len(array) - 1; i >= 0; i-- {
		temp = append(temp, array[i])
	}

	return temp
}

func swap(source, temp []byte, start, length int) {

	for i := 0; i < len(temp); i++ {
		source[(start+i)%length] = temp[i]
	}
}

func encrypt(array, inputs []byte) string {
	skip := 0
	p := 0
	var stringBuf bytes.Buffer

	for x := 1; x <= 64; x++ {
		p, skip = round(array, inputs, p, skip)
	}
	for x := 0; x < 256; x += 16 {
		stringBuf.WriteString(xorHex(array[x : x+15]))
	}

	return stringBuf.String()
}

func round(array, inputs []byte, p, skip int) (int, int) {
	sourceLength := len(array)
	for _, i := range inputs {

		input := int(i)

		sub := slice(array, p, input, sourceLength)

		rev := reverse(sub)

		swap(array, rev, p, sourceLength)
		p = (p + input + skip) % sourceLength
		skip++

	}

	return p, skip
}

func xorHex(slice []byte) string {
	xor := int(slice[0])
	for _, value := range slice[1:16] {
		xor = xor ^ int(value)
	}
	return strconv.FormatInt(int64(xor), 16)

}

func slice(source []byte, start, stop, sourceLength int) []byte {
	var temp []byte
	startIndex := int(start)
	for i := 0; i < stop; i++ {
		temp = append(temp, source[(i+startIndex)%sourceLength])
	}
	return temp
}
