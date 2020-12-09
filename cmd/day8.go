package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type BootLine struct {
	command string
	value   int
}

func (code *BootLine) switchCommand(first string, second string) bool {
	if code.command == first {
		code.command = second
		return true
	} else if code.command == second {
		code.command = first
		return true
	} else {
		return false
	}
}

func (code *BootLine) setFromString(line string) {
	code.command = line[:3]
	_step, err := strconv.ParseInt(line[5:], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	step := int(_step)
	if line[4] == '-' {
		step = -step
	}
	code.value = step
}

type BootCode struct {
	code []BootLine
}

func (boot *BootCode) switchCommand(first string, second string, index int) bool {
	return boot.code[index].switchCommand(first, second)
}

func (boot *BootCode) execute() (int, bool) {
	var visited []int
	var position, output int

	for position < len(boot.code) {
		visited = append(visited, position)
		switch boot.code[position].command {
		case "nop":
			position += 1
		case "acc":
			output += boot.code[position].value
			position += 1
		case "jmp":
			position += boot.code[position].value
		}
		if containsInt(visited, position) {
			return output, false
		}
	}
	return output, true
}

func (boot *BootCode) setFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newLine := new(BootLine)
		newLine.setFromString(scanner.Text())
		boot.code = append(boot.code, *newLine)
	}
}

func forceRunCode(filename string) int {
	var boot BootCode
	boot.setFromFile(filename)
	output, _ := boot.execute()
	return output
}

func debugCode(filename string) int {
	var boot BootCode
	boot.setFromFile(filename)
	for ii := 0; ii < len(boot.code); ii++ {
		switched := boot.switchCommand("nop", "jmp", ii)
		if !switched {
			continue
		}
		output, success := boot.execute()
		if success {
			return output
		}
		boot.switchCommand("nop", "jmp", ii)
	}
	return -1
}

func Day8() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", forceRunCode("inputs/day8_test.txt"))
	fmt.Println("Part 2: ", debugCode("inputs/day8_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", forceRunCode("inputs/day8.txt"))
	fmt.Println("Part 2: ", debugCode("inputs/day8.txt"))
}
