package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Read an array of integers from a file
// The first argument is the file name
// The second argument is the delimiter, e.g. "\n", ","
func readIntFile(inputFile string, sep string) (numbers []int, err error) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(data), sep)
	numbers = make([]int, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 { continue }
		n, err := strconv.Atoi(line)
		if err != nil { return nil, err }
		numbers = append(numbers, n)
	}
	return numbers, nil
}

// Find the first pair of elements which sum to 2020 and
// return the product of those two numbers.
func findPairSum(data []int) int {
	for ii, value1 := range data {
		for _, value2 := range data[ii:] {
			if (value1 + value2) == 2020 {
				return value1 * value2
			}
		}
	}
	return -1
}


// Find the trio of elements which sum to 2020 and
// return the product of those three numbers.
func findTrioSum(data []int) int {
	for ii, value1 := range data {
		for jj, value2 := range data[ii:] {
			if (value1 + value2) < 2020 {
				for _, value3 := range data[jj:] {
					if (value1 + value2 + value3) == 2020 {
						return value1 * value2 * value3
					}
				}
			}
		}
	}
	return -1
}


func main() {
	inputFile := "inputs/day1.txt"
	data, err := readIntFile(inputFile, "\n")
	if err != nil {fmt.Println(err)}

	fmt.Println("=========")
	fmt.Println("Part 1")
	fmt.Println("=========")
	fmt.Println(findPairSum(data))
	fmt.Println("=========")
	fmt.Println("Part 2")
	fmt.Println("=========")
	fmt.Println(findTrioSum(data))
	fmt.Println("=========")
}