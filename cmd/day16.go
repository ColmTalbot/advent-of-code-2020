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

func validateTicketEntry(rule [][]int, value int) bool {
	for _, pair := range rule {
		if value >= pair[0] && value <= pair[1] {
			return true
		}
	}
	return false
}

func validateTicket(rules [][][]int, ticket []int) (output bool) {
	output = true
	for _, entry := range ticket {
		valid := false
		for _, rule := range rules {
			valid = valid || validateTicketEntry(rule, entry)
		}
		output = output && valid
	}
	return
}

func parseTicketFile(filename string) (rules [][][]int, tickets [][]int, labels []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+-[0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Split(line, "or")) > 1 {
			labels = append(labels, strings.Split(line, ":")[0])
			var _rule [][]int
			for _, rule := range re.FindAllString(line, -1) {
				bounds := strings.Split(rule, "-")
				low, err := strconv.Atoi(bounds[0])
				if err != nil {
					log.Fatal(err)
				}
				high, err := strconv.Atoi(bounds[1])
				if err != nil {
					log.Fatal(err)
				}
				_rule = append(_rule, []int{low, high})
			}
			rules = append(rules, _rule)
		} else if len(strings.Split(line, ",")) > 1 {
			var ticket []int
			for _, entry := range strings.Split(line, ",") {
				value, err := strconv.Atoi(entry)
				if err != nil {
					log.Fatal(err)
				}
				ticket = append(ticket, value)
			}
			tickets = append(tickets, ticket)
		}
	}
	return
}

func scanningErrorRate(filename string) int {
	rules, tickets, _ := parseTicketFile(filename)
	output := 0
	for _, ticket := range tickets[1:] {
		for _, entry := range ticket {
			valid := false
			for _, rule := range rules {
				valid = valid || validateTicketEntry(rule, entry)
			}
			if !valid {
				output += entry
			}
		}
	}
	return output
}

// First remove the bad tickets
// Next find all possible rule IDs for each column
// Then assume that one column will allow one rule
// one column will allow two rules, etc.
// This is true to the input data.
// Sort the possible IDs by the number of allowed rules
// then loop one last time over the increasing number of rules
func decodeTicket(filename string) int {
	rules, tickets, labels := parseTicketFile(filename)
	var validTickets [][]int
	for _, ticket := range tickets[1:] {
		if validateTicket(rules, ticket) {
			validTickets = append(validTickets, ticket)
		}
	}

	var label string
	var identified []int
	allPossible := make([][]int, len(tickets[0]))
	ordering := make([]int, len(tickets[0]))
	for kk := range tickets[0] {
		var possible []int
		for jj, rule := range rules {
			valid := true
			for ii := range validTickets {
				valid = valid && validateTicketEntry(rule, validTickets[ii][kk])
			}
			if valid {
				possible = append(possible, jj)
			}
		}
		ordering[len(possible)-1] = kk
		allPossible[kk] = possible
	}

	output := 1
	for _, kk := range ordering {
		for _, jj := range allPossible[kk] {
			if containsInt(identified, jj) {
				continue
			}
			label = labels[jj]
			identified = append(identified, jj)
			if strings.HasPrefix(label, "departure") {
				output *= tickets[0][kk]
			}
		}
	}

	return output
}

func Day16() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", scanningErrorRate("inputs/day16_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", scanningErrorRate("inputs/day16.txt"))
	fmt.Println("Part 2: ", decodeTicket("inputs/day16.txt"))
}
