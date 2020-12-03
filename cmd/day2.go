package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Parse a line with format 1-2 a: abcde
// Returns the two integers, the single letter and the string
func parsePasswordLine(line string) (first int, second int, target string, password string, err error) {
	line = strings.Replace(line, ":", "", 1)
	splitLine := strings.Split(line, " ")
	first, err = strconv.Atoi(strings.Split(splitLine[0], "-")[0])
	if err != nil { return 0, 0, "", "", err }
	second, err = strconv.Atoi(strings.Split(splitLine[0], "-")[1])
	if err != nil { return 0, 0, "", "", err }
	target = splitLine[1]
	password = splitLine[2]
	return
}

// Verify that the number of occurrences of the target letter is between the maximum
// and minimum.
func checkCountAllowed(password string, target string, min int, max int) (allowed bool) {
	occurrences := strings.Count(password, target)
	allowed = min <= occurrences && max >= occurrences
	return
}

// Check that the target value is present at only one of the two provided indices
func checkPositionAllowed(password string, target string, first int, second int) (allowed bool) {
	firstAllowed := string(password[first - 1]) == target
	secondAllowed := string(password[second - 1]) == target
	allowed = (firstAllowed || secondAllowed) && !(firstAllowed && secondAllowed)
	return
}

// Find the valid passwords stored in the file using the two criteria above
func findValidPasswords(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nAllowed1 := 0
	nAllowed2 := 0
	for scanner.Scan() {
		first, second, target, password, err := parsePasswordLine(scanner.Text())
		if err != nil { fmt.Println(err) }
		if checkCountAllowed(password, target, first, second) {
			nAllowed1 += 1
		}
		if checkPositionAllowed(password, target, first, second) {
			nAllowed2 += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", nAllowed1)
	fmt.Println("Part 2:", nAllowed2)

}

/*
--- Day 2: Password Philosophy ---
Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we can't log in!" You ask if you can take a look.

Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy that was in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?

--- Part Two ---
While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.

Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
How many passwords are valid according to the new interpretation of the policies?
*/
func Day2() {
	fmt.Println("Test")
	findValidPasswords("inputs/day2_test.txt")
	fmt.Println("Main")
	findValidPasswords("inputs/day2.txt")
}