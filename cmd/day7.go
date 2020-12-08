package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseBagFile(filename string) (map[string][]string, map[string]map[string]int64) {
	var inverseMap map[string][]string
	inverseMap = make(map[string][]string)
	var valueMap map[string]map[string]int64
	valueMap = make(map[string]map[string]int64)
	var temp []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp = parseBagLine(scanner.Text())
		valueMap[temp[0]] = make(map[string]int64)
		for _, color := range temp[1:] {
			inverseMap[color[1:]] = append(inverseMap[color[1:]], temp[0])
			if color == "0other" {
				continue
			}
			valueMap[temp[0]][color[1:]], err = strconv.ParseInt(string(color[0]), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return inverseMap, valueMap
}

func parseBagLine(line string) []string {
	line = strings.ReplaceAll(line, "bags", "")
	line = strings.ReplaceAll(line, "bag", "")
	line = strings.ReplaceAll(line, ".", "")
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "contain", ",")
	line = strings.ReplaceAll(line, "no", "0")
	return strings.Split(line, ",")
}

func containingBags(mapping map[string][]string, target string, depth int) (containers []string) {
	for _, color := range mapping[target] {
		if !containsString(containers, color) {
			containers = append(containers, color)
		}
		for _, temp := range containingBags(mapping, color, depth+1) {
			if !containsString(containers, temp) {
				containers = append(containers, temp)
			}
		}
	}
	return
}

func containedBags(mapping map[string][]string, values map[string]map[string]int64, target string, depth int) int64 {
	var total int64
	if len(mapping[target]) == 0 {
		total = 1
	} else {
		total = 0
	}
	for color := range values[target] {
		total += values[target][color] * (1 + containedBags(mapping, values, color, depth+1))
	}
	return total
}

func goldContainingBags(filename string) int {
	inverseMapping, _ := parseBagFile(filename)
	allBags := containingBags(inverseMapping, "shinygold", 0)
	var output []string
	total := 0
	for _, bag := range allBags {
		if !containsString(output, bag) {
			total += 1
			output = append(output, bag)
		}
	}
	return total
}

func countGoldContainingBags(filename string) (total int64) {
	reverseMapping, valueMapping := parseBagFile(filename)
	return containedBags(reverseMapping, valueMapping, "shinygold", 0)
}

func Day7() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", goldContainingBags("inputs/day7_test.txt"))
	fmt.Println("Part 2: ", countGoldContainingBags("inputs/day7_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", goldContainingBags("inputs/day7.txt"))
	fmt.Println("Part 2: ", countGoldContainingBags("inputs/day7.txt"))
}
