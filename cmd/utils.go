package advent

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseSetOfLines(filename string) (items []string) {
	allData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	items = strings.Split(string(allData), "\n\n")
	return
}

func containsString(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func containsInt(slice []int, target int) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

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
		if len(line) == 0 {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
