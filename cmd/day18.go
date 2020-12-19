package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evaluateImmediateString(input string) (output int) {
	entries := strings.Split(input, " ")
	output = parseInt(entries[0], 10)
	operand := ""
	for _, entry := range entries[1:] {
		if entry == "+" || entry == "*" {
			operand = entry
		} else if operand == "+" {
			output += parseInt(entry, 10)
		} else if operand == "*" {
			output *= parseInt(entry, 10)
		}
	}
	return
}

func replaceParentheses(line string) string {
	re := regexp.MustCompile("\\([0-9 +*]+\\)")
	for _, subline := range re.FindAllString(line, -1) {
		line = strings.Replace(line, subline, strconv.Itoa(evaluateImmediateString(subline[1:len(subline)-1])), 1)
	}
	return line
}

func replaceParenthesesAddition(line string) string {
	re := regexp.MustCompile("\\([0-9 +*]+\\)")
	for _, subline := range re.FindAllString(line, -1) {
		line = strings.Replace(line, subline, strconv.Itoa(evaluateImmediateString(performAddition(subline[1:len(subline)-1]))), 1)
	}
	return line
}

func performAddition(line string) string {
	re := regexp.MustCompile("[0-9]+ \\+ [0-9]+")
	for strings.Contains(line, "+") {
		subline := re.FindString(line)
		line = strings.Replace(line, subline, strconv.Itoa(evaluateImmediateString(subline)), 1)
	}
	return line
}

func evaluateImmediate(filename string) (output int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for strings.Contains(line, "(") {
			line = replaceParentheses(line)
		}
		output += evaluateImmediateString(line)
	}
	return output
}

func evaluateAdditionFirst(filename string) (output int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for strings.Contains(line, "(") {
			line = replaceParenthesesAddition(line)
		}
		line = performAddition(line)
		output += evaluateImmediateString(line)
	}
	return output
}

func Day18() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", evaluateImmediate("inputs/day18_test.txt"))
	fmt.Println("Part 2: ", evaluateAdditionFirst("inputs/day18_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", evaluateImmediate("inputs/day18.txt"))
	fmt.Println("Part 2: ", evaluateAdditionFirst("inputs/day18.txt"))
}
