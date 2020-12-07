package advent

import (
	"io/ioutil"
	"log"
	"strings"
)

func parseSetOfLines(filename string) (items []string) {
	allData, err := ioutil.ReadFile(filename)
	if err != nil { log.Fatal(err) }

	items = strings.Split(string(allData), "\n\n")
	return
}

