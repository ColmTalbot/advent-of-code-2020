package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parsePassportFile(filename string) (passports []string) {
	allPassports, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	passports = strings.Split(string(allPassports), "\n\n")
	return
}

// Parse passport string to check if required fields are present
// If any required fields are missing return 0
// Else return 1
func assessPassportFields(passport string) bool {
	requiredFields := []string{"ecl:", "pid:", "eyr:", "hcl:", "byr:", "iyr:", "hgt:"}
	for _, field := range requiredFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func isDelimiter(r rune) bool {
	return r == ' ' || r == '\n'
}

func checkIntegerInRange(input string, minimum int, maximum int) bool {
	value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if value < minimum || value > maximum {
		return false
	} else {
		return true
	}
}

func checkEyeColour(input string) bool {
	switch input {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func checkHeightInRange(input string) bool {
	if !(len(input) == 4 || len(input) == 5) {
		return false
	}
	if input[len(input)-2:] == "cm" {
		return checkIntegerInRange(input[:len(input)-2], 150, 193)
	} else if input[len(input)-2:] == "in" {
		return checkIntegerInRange(input[:len(input)-2], 59, 76)
	} else {
		return false
	}
}

func assessPassportValidity(passport string) bool {
	if !assessPassportFields(passport) {
		return false
	}

	var pair []string

	passportFields := strings.FieldsFunc(passport, isDelimiter)
	for _, field := range passportFields {
		pair = strings.Split(field, ":")
		switch pair[0] {
		case "byr":
			if !checkIntegerInRange(pair[1], 1920, 2002) {
				return false
			}
		case "iyr":
			if !checkIntegerInRange(pair[1], 2010, 2020) {
				return false
			}
		case "eyr":
			if !checkIntegerInRange(pair[1], 2020, 2030) {
				return false
			}
		case "hgt":
			if !checkHeightInRange(pair[1]) {
				return false
			}
		case "hcl":
			// regex: match something like #123abc
			// "#" followed by exactly 6 hexadecimal characters
			if !regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(pair[1]) {
				return false
			}
		case "ecl":
			if !checkEyeColour(pair[1]) {
				return false
			}
		case "pid":
			// regex: match something like 012345678
			// exactly 9 integers
			if !regexp.MustCompile(`^\d{9}$`).MatchString(pair[1]) {
				return false
			}
		}
	}
	return true
}

func findValidPassports(filename string, strict bool) (validPassports []string) {
	passports := parsePassportFile(filename)
	var condition func(string) bool
	if strict {
		condition = assessPassportValidity
	} else {
		condition = assessPassportFields
	}
	for _, passport := range passports {
		if condition(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return
}

func countValidPassports(filename string, strict bool) int {
	passports := findValidPassports(filename, strict)
	return len(passports)
}

func Day4() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", countValidPassports("inputs/day4_test.txt", false))
	fmt.Println("Part 2: ", countValidPassports("inputs/day4_test.txt", true))
	fmt.Println("Main")
	fmt.Println("Part 1: ", countValidPassports("inputs/day4.txt", false))
	fmt.Println("Part 2: ", countValidPassports("inputs/day4.txt", true))
}
