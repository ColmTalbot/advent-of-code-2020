package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func getMemoryValue(line string, mask string) (output int) {
	for ii, character := range mask {
		output = output << 1
		if character == '1' {
			output++
		} else if character == '0' {
		} else if line[ii] == '1' {
			output++
		}
	}
	return
}

func getMemoryAddresses(line string, mask string) (addresses []string) {
	var newLine string
	for ii := 0; ii < len(mask); ii++ {
		if mask[ii] == '1' {
			newLine += "1"
		} else if mask[ii] == '0' {
			newLine += string(line[ii])
		} else {
			for _, address := range getMemoryAddresses(line[ii+1:], mask[ii+1:]) {
				if len(newLine)+len(address)+1 == len(line) {
					addresses = append(addresses, newLine+"0"+address)
					addresses = append(addresses, newLine+"1"+address)
				}
			}
		}
	}
	if len(newLine) == len(mask) {
		addresses = append(addresses, newLine)
	}
	return
}

func formatValue(input string, length int) (output string) {
	_value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	output = strconv.FormatInt(int64(_value), 2)
	for len(output) < length {
		output = "0" + output
	}
	return
}

func memorySum(filename string, part int) int {
	_data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(_data), "\n")
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	var values []string
	memory := make(map[string]int)
	re := regexp.MustCompile("[0-9]+")

	for _, line := range data {
		if line[:4] == "mask" {
			mask = line[7:]
		} else {
			values = re.FindAllString(line, -1)
			if part == 1 {
				memory[values[0]] = getMemoryValue(formatValue(values[1], len(mask)), mask)
			} else {
				value, err := strconv.Atoi(values[1])
				if err != nil {
					log.Fatal(err)
				}
				for _, address := range getMemoryAddresses(formatValue(values[0], len(mask)), mask) {
					memory[address] = value
				}
			}
		}
	}
	total := 0
	for _, val := range memory {
		if part == 2 {
			total += val
		} else {
			total += val
		}
	}
	return total
}

func Day14() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", memorySum("inputs/day14_test.txt", 1))
	fmt.Println("Part 2: ", memorySum("inputs/day14_test.txt", 2))
	fmt.Println("Main")
	fmt.Println("Part 1: ", memorySum("inputs/day14.txt", 1))
	fmt.Println("Part 2: ", memorySum("inputs/day14.txt", 2))
}
