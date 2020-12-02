package main

import (
	advent "advent-of-code-2020/cmd"
	"fmt"
	"strings"
)

func main() {
	var days []func()
	days = append(days, advent.Day1)
	lineBreak := strings.Repeat("=", 20)
	for ii, function := range days {
		fmt.Println(lineBreak)
		fmt.Println("Day ", ii+1)
		fmt.Println(lineBreak)
		function()
	}
	fmt.Println(lineBreak)
}
