package main

import (
	advent "advent-of-code-2020/cmd"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Run the advent of code challenges for 2020.
// A specific day or range of days can be run using command-line arguments
// go run advent.go 2   : runs just day 2
// go run advent.go 1 3 : runs just days 1-3
func main() {
	var days []func()
	var err error
	days = append(days, advent.Day1)
	days = append(days, advent.Day2)
	days = append(days, advent.Day3)
	days = append(days, advent.Day4)
	days = append(days, advent.Day5)
	days = append(days, advent.Day6)
	days = append(days, advent.Day7)
	days = append(days, advent.Day8)
	days = append(days, advent.Day9)
	days = append(days, advent.Day10)
	days = append(days, advent.Day11)
	days = append(days, advent.Day12)
	days = append(days, advent.Day13)
	days = append(days, advent.Day14)

	lineBreak := strings.Repeat("=", 30)

	start := 1
	end := len(days)
	if len(os.Args) == 2 {
		start, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		end = start
	} else if len(os.Args) == 3 {
		start, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		end, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
	}

	for ii, function := range days[start-1 : end] {
		fmt.Println(lineBreak)
		fmt.Println("Day", ii+start)
		fmt.Println(lineBreak)
		startTime := time.Now()
		function()
		duration := time.Since(startTime)
		fmt.Println(lineBreak)
		fmt.Println("Evaluation time: ", duration)
	}
	fmt.Println(lineBreak)
}
