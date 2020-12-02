/*
--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?

--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?
*/
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