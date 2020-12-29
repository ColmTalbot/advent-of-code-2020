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

func readMessages(filename string) (rules map[int][]string, messages []string, known []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inMessages := false
	rules = make(map[int][]string)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "" {
			inMessages = true
		} else if inMessages {
			messages = append(messages, message)
		} else {
			content := strings.Split(message, ":")
			index, err := strconv.Atoi(content[0])
			if err != nil {
				log.Fatal(err)
			}
			if strings.ContainsRune(message, '"') {
				rules[index] = append(rules[index], content[1][2:3])
				known = append(known, index)
			} else {
				for _, temp := range strings.Split(content[1][1:], " | ") {
					rules[index] = append(rules[index], temp)
				}
			}
		}
	}
	return
}

func nPossibleMessages(filename string, ruleIdx int) (output int) {
	_, messages, _ := readMessages(filename)
	compiledRules := compileRules(filename)
	targets := compiledRules[ruleIdx]
	for _, message := range messages {
		if containsString(targets, message) {
			output += 1
		}
	}
	return
}

func recurseLevels(levels map[int]int, rules map[int][]string) {
	re := regexp.MustCompile("[0-9]+")
	for key, rule := range rules {
		level := 1
		failed := false
		for _, subrule := range rule {
			for _, value := range re.FindAllString(subrule, -1) {
				temp, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				if levels[temp] == 0 {
					failed = true
				}
				if levels[temp] >= level {
					level = levels[temp] + 1
				}
			}
			if failed {
				continue
			}
			levels[key] = level
		}
	}
}

func computeLevels(filename string) (levels map[int]int) {
	rules, _, known := readMessages(filename)
	levels = make(map[int]int)
	for _, key := range known {
		levels[key] = 1
	}
	oldLength := 0
	for oldLength != len(levels) {
		oldLength = len(levels)
		recurseLevels(levels, rules)
	}
	return
}

func rulesPerLevel(filename string) (levelRules map[int][]int) {
	levels := computeLevels(filename)
	levelRules = make(map[int][]int)
	for rule, level := range levels {
		levelRules[level] = append(levelRules[level], rule)
	}
	return
}

func compileRule(rule string, rules map[int][]string) (compiledRule []string) {
	re := regexp.MustCompile("[0-9]+")
	nRules := 1
	for _, value := range re.FindAllString(rule, -1) {
		idx, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		nRules *= len(rules[idx])
		if len(compiledRule) == 0 {
			compiledRule = rules[idx]
		} else {
			temp := make([]string, len(compiledRule)*len(rules[idx]))
			for ii, rule1 := range compiledRule {
				for jj, rule2 := range rules[idx] {
					temp[ii*len(rules[idx])+jj] = rule1 + rule2
				}
			}
			compiledRule = temp
		}
	}
	return
}

func compileRules(filename string) (compiledRules map[int][]string) {
	rules, _, _ := readMessages(filename)
	levelRules := rulesPerLevel(filename)
	compiledRules = make(map[int][]string)
	maxLevel := 0
	maxLength := 88
	for level := range levelRules {
		if level > maxLevel {
			maxLevel = level
		}
	}
	for level := 1; level <= maxLevel; level++ {
		_rules := levelRules[level]
		if level == 1 {
			for _, rule := range _rules {
				compiledRules[rule] = rules[rule]
			}
		} else {
			for _, rule := range _rules {
				var newRule []string
				if rule == -8 {
					_rule := compileRule(rules[rule][0], compiledRules)
					for _, subrule := range _rule {
						temp := subrule
						for len(temp) < maxLength {
							if !containsString(newRule, temp) {
								newRule = append(newRule, temp)
							}
							temp = temp + subrule
						}
					}
					fmt.Println(newRule)
				} else {
					for _, subrule := range rules[rule] {
						for _, temp := range compileRule(subrule, compiledRules) {
							newRule = append(newRule, temp)
						}
					}
				}
				compiledRules[rule] = newRule
			}
		}
	}
	return
}

func nPossibleMessages2(filename string, ruleIdx int) (output int) {
	_, messages, _ := readMessages(filename)
	compiledRules := compileRules(filename)
	targets := compiledRules[ruleIdx]
	subLength := len(compiledRules[42][0])
	for _, message := range messages {
		if containsString(targets, message) {
			output++
		} else if len(message)%subLength == 0 {
			nInitialMatch, nFinalMatch := 0, 0
			for ii := 0; ii < len(message)-subLength; ii += subLength {
				if containsString(compiledRules[42], message[ii:ii+subLength]) {
					nInitialMatch += 1
				} else {
					break
				}
			}
			for ii := len(message) - subLength; ii >= nInitialMatch*subLength; ii -= subLength {
				if containsString(compiledRules[31], message[ii:ii+subLength]) {
					nFinalMatch += 1
				} else {
					break
				}
			}
			if nInitialMatch+nFinalMatch == len(message)/subLength && nInitialMatch > nFinalMatch && nFinalMatch > 0 {
				output++
			}
		}
	}
	return
}

// for each instruction list explicit dependencies (variables)
// if instruction has no dependencies assign level 0
// if all dependencies are level 0, assign level 1
// if all dependencies are level n or below, assign level n + 1
// in example:
// - level 0: 4, 5
// - level 1: 2, 3
// - level 2: 1
// - level 3: 0
//
// compile level n instructions starting with zero
// descend down levels looking for a match for each part
//
// does each instruction output have a unique length?
// each subinstruction does have unique length
// use this to determine spacing for three part instructions
//
// Part 2: add extra rules
// 8: can be any combination of rules satisfying 8
// 11: 42 11 31 becomes 42 42 ... 42 31 ... 31 31
// Total pattern: 42 * (m + n) + 31 * n
func Day19() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", nPossibleMessages("inputs/day19_test.txt", 0))
	fmt.Println("Part 2: ", nPossibleMessages2("inputs/day19_test.txt", 0))
	fmt.Println("Main")
	fmt.Println("Part 1: ", nPossibleMessages("inputs/day19.txt", 0))
	fmt.Println("Part 2: ", nPossibleMessages2("inputs/day19.txt", 0))
}
