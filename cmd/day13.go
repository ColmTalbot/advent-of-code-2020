package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func waitTime(filename string) int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	splitData := strings.Split(string(data), "\n")

	earliest, err := strconv.Atoi(splitData[0])
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("[0-9]+")
	inputs := re.FindAllString(splitData[1], -1)
	var output, number, wait int
	minimumWait := 100000000000

	for _, value := range inputs {
		number, err = strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		wait = ((earliest/number + 1) * number) - earliest
		if wait < minimumWait {
			minimumWait = wait
			output = wait * number
		}
	}
	return output
}

func linearLeaving(filename string) (output int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	splitData := strings.Split(string(data), "\n")[1]
	inputs := strings.Split(splitData, ",")
	numbers := make([]int, len(inputs))

	for ii, input := range inputs {
		if input == "x" {
			numbers[ii] = 0
		} else {
			value, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal(err)
			}
			numbers[ii] = value
		}
	}
	output = numbers[0]
	step := numbers[0]
	for ii, value := range numbers[1:] {
		if value == 0 {
			continue
		}
		for (output/value+1)*value-output != (ii+1)%value {
			output += step
		}
		step *= value / gcd(step, value)
	}
	return
}

func Day13() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", waitTime("inputs/day13_test.txt"))
	fmt.Println("Part 2: ", linearLeaving("inputs/day13_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", waitTime("inputs/day13.txt"))
	fmt.Println("Part 2: ", linearLeaving("inputs/day13.txt"))
}
