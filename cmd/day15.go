package advent

import (
	"fmt"
	"log"
)

func nthNumber(filename string, nn int) int {
	data, err := readIntFile(filename, ",")
	if err != nil {
		log.Fatal(err)
	}
	past := make([]int, nn)
	counter := 0
	temp := 0
	value := 0
	past[0] = 1

	for _, entry := range data {
		past[value] = counter
		counter++
		value = entry
	}

	for counter < nn {
		temp = past[value]
		past[value] = counter
		if temp > 0 {
			value = counter - temp
		} else {
			value = 0
		}
		counter++
	}

	return value
}

func Day15() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", nthNumber("inputs/day15_test.txt", 2020))
	fmt.Println("Part 2: ", nthNumber("inputs/day15_test.txt", 30000000))
	fmt.Println("Main")
	fmt.Println("Part 1: ", nthNumber("inputs/day15.txt", 2020))
	fmt.Println("Part 2: ", nthNumber("inputs/day15.txt", 30000000))
}
