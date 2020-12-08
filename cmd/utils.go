package advent

import (
	"io/ioutil"
	"log"
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
