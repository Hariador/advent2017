package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Defrag")
	argCount := len(os.Args)
	if argCount < 2 {
		panic("day14 <key>")
	}

	key := os.Args[1]

	//array := initArray(128)
	keys := getSourceData(key)
	printKeys(keys)

}

func initArray(size int) []byte {
	array := make([]byte, size)
	for k := range array {
		array[k] = byte(k)
	}
	return array

}

func getSourceData(key string) [][]byte {
	var keys [][]byte

	dash := byte('-')
	source := []byte(key)
	var disk [][]sector
	disk = make([][]sector, 128)
	for i := range disk {
		disk[i] = make([]sector, 128)
	}
	for x := 0; x < 128; x++ {
		xString := []byte(strconv.Itoa(x))
		temp := make([]byte, 0)
		temp = append(temp, source[:]...)
		temp = append(temp, dash)
		temp = append(temp, xString[:]...)
		temp = append(temp, 17, 31, 73, 47, 23)

		result := []byte(encrypt(initArray(256), temp))

		if len(result) != 128 {
			panic("OH FUCK")
		}
		//	total = total + countHash(result)

		initDisk(disk, result, x)

	}
	group := 1
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if disk[x][y].Check() {
				disk[x][y].Scan(disk, group)
				group++

			}
		}
	}
	fmt.Println(group - 1)
	printDisk(disk)
	return keys

}

func printKeys(keys [][]byte) {
	for _, key := range keys {
		fmt.Println(key)
	}
}

func printDisk(disk [][]sector) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if disk[i][j].full {
				fmt.Print(disk[i][j].groupID)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}

}

func initDisk(disk [][]sector, hash []byte, row int) {
	for i := 0; i < 128; i++ {
		if hash[i]-48 > 0 {
			disk[row][i].Mark(row, i)
		}
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
		stringBuf.WriteString(xorBin(array[x : x+15]))
	}

	return stringBuf.String()
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

func xorBin(slice []byte) string {
	xor := int(slice[0])
	for _, value := range slice[1:16] {
		xor = xor ^ int(value)
	}

	binString := strconv.FormatInt(int64(xor), 2)
	diff := 8 - len(binString)
	for i := 0; i < diff; i++ {
		binString = "0" + binString
	}

	return binString
}

func slice(source []byte, start, stop, sourceLength int) []byte {
	var temp []byte
	startIndex := int(start)
	for i := 0; i < stop; i++ {
		temp = append(temp, source[(i+startIndex)%sourceLength])
	}
	return temp
}

type sector struct {
	full    bool
	groupID int
	x       int
	y       int
}

func (s *sector) Mark(x, y int) {
	s.full = true
	s.x = x
	s.y = y
}

func (s *sector) Scan(disk [][]sector, group int) int {
	groupSize := 0
	s.groupID = group
	if s.x-1 >= 0 {
		if disk[s.x-1][s.y].Check() {
			groupSize += disk[s.x-1][s.y].Scan(disk, group)
		}
	}

	if s.x+1 < 128 {
		if disk[s.x+1][s.y].Check() {
			groupSize += disk[s.x+1][s.y].Scan(disk, group)
		}
	}

	if s.y-1 >= 0 {
		if disk[s.x][s.y-1].Check() {
			groupSize += disk[s.x][s.y-1].Scan(disk, group)
		}
	}

	if s.y+1 < 128 {
		if disk[s.x][s.y+1].Check() {
			groupSize += disk[s.x][s.y+1].Scan(disk, group)
		}
	}

	return groupSize

}

func (s *sector) Check() bool {
	return s.groupID == 0 && s.full
}
