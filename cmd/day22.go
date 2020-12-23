package advent

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func loadHands(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	player := 0
	hands := make([][]int, 2)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			player += 1
		} else if line[0] == 'P' {
			continue
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			hands[player] = append(hands[player], value)
		}
	}
	return hands
}

func scoreHand(hand []int) (output int) {
	for ii, value := range hand {
		output += value * (len(hand) - ii)
	}
	return
}

func combatGame(hands [][]int) []int {
	var left, right, winner int
	for len(hands[0]) > 0 && len(hands[1]) > 0 {
		left, hands[0] = hands[0][0], hands[0][1:]
		right, hands[1] = hands[1][0], hands[1][1:]
		if left > right {
			winner = 0
			hands[0] = append(hands[0], left)
			hands[0] = append(hands[0], right)
		} else {
			winner = 1
			hands[1] = append(hands[1], right)
			hands[1] = append(hands[1], left)
		}
	}
	return hands[winner]
}

func playCombat(filename string) (output int) {
	hands := loadHands(filename)
	winner := combatGame(hands)
	output = scoreHand(winner)
	return
}

func copyHand(hand []int, length int) []int {
	output := make([]int, length)
	for ii, value := range hand[:length] {
		output[ii] = value
	}
	return output
}

func copyHands(hands [][]int, left, right int) [][]int {
	output := make([][]int, 2)
	output[0] = copyHand(hands[0], left)
	output[1] = copyHand(hands[1], right)
	return output
}

func computeHashKeyForHands(hands [][]int) (hash string) {
	hash += strconv.FormatInt(int64(len(hands[0])), 10)
	hash += strconv.FormatInt(int64(len(hands[1])), 10)
	for _, hand := range hands {
		hash += "xx"
		for _, value := range hand {
			hash += strconv.FormatInt(int64(value), 10)
		}
	}
	return
}

func recursiveCombat(hands [][]int) (winner int) {
	var left, right int
	var hash string
	var previousRounds []string
	max1 := float64(int64Max(hands[0]))
	max2 := float64(int64Max(hands[1]))
	globalMax := math.Max(max1, max2)
	// If player 1 has the largest value which is larger
	// than the total number of cards in play, player 1
	// wins.
	if int(globalMax) > len(hands[0])+len(hands[1]) {
		if max1 > max2 {
			return 0
		}
	}
	for len(hands[0]) > 0 && len(hands[1]) > 0 {
		hash = computeHashKeyForHands(hands)
		left, hands[0] = hands[0][0], hands[0][1:]
		right, hands[1] = hands[1][0], hands[1][1:]
		if containsString(previousRounds, hash) {
			return 0
		} else if len(hands[0]) >= left && len(hands[1]) >= right {
			newHands := copyHands(hands, left, right)
			winner = recursiveCombat([][]int{newHands[0][:left], newHands[1][:right]})
			if winner == 0 {
				hands[0] = append(hands[0], left)
				hands[0] = append(hands[0], right)
			} else {
				hands[1] = append(hands[1], right)
				hands[1] = append(hands[1], left)
			}
		} else if left > right {
			winner = 0
			hands[0] = append(hands[0], left)
			hands[0] = append(hands[0], right)
		} else {
			winner = 1
			hands[1] = append(hands[1], right)
			hands[1] = append(hands[1], left)
		}
		previousRounds = append(previousRounds, hash)
	}
	return
}

func playRecursiveCombat(filename string) (output int) {
	hands := loadHands(filename)
	winner := recursiveCombat(hands)
	output = scoreHand(hands[winner])
	return
}

func Day22() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", playCombat("inputs/day22_test.txt"))
	fmt.Println("Part 2: ", playRecursiveCombat("inputs/day22_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", playCombat("inputs/day22.txt"))
	fmt.Println("Part 2: ", playRecursiveCombat("inputs/day22.txt"))
}
