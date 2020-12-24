package advent

import (
	"container/ring"
	"fmt"
	"log"
)

func naturalMod(value, base int) int {
	return (value+base-1)%base + 1
}

func iterateCups(data *ring.Ring, positions map[int]*ring.Ring, length int, iterations int) *ring.Ring {
	values := make([]int, 3)
	var target int
	var excludes *ring.Ring
	for current := 0; current < iterations; current++ {
		target = naturalMod(data.Value.(int)-1, length)
		excludes = data.Unlink(3)
		for ii := 0; ii < 3; ii++ {
			values[ii] = excludes.Value.(int)
			excludes = excludes.Next()
		}
		for containsInt(values, target) {
			target = naturalMod(target-1, length)
		}
		positions[target].Link(excludes)
		data = data.Next()
	}
	return data
}

func finalCups(cups *ring.Ring) (output int) {
	for ii := 1; ii < 9; ii++ {
		cups = cups.Next()
		output *= 10
		output += cups.Value.(int)
	}
	return
}

func loadCups(filename string, length int) (cups *ring.Ring, positions map[int]*ring.Ring) {
	var cup int
	cups = ring.New(length)
	positions = make(map[int]*ring.Ring)
	data, err := readIntFile(filename, "")
	if err != nil {
		log.Fatal(err)
	}
	for ii := 0; ii < length; ii++ {
		if ii < len(data) {
			cup = data[ii]
		} else {
			cup = ii + 1
		}
		cups.Value = cup
		positions[cup] = cups
		cups = cups.Next()
	}
	return
}

func cups1(filename string) int {
	cups, positions := loadCups(filename, 9)
	cups = iterateCups(cups, positions, 9, 100)
	return finalCups(positions[1])
}

func cups2(filename string) int {
	cups, positions := loadCups(filename, 1000000)
	cups = iterateCups(cups, positions, 1000000, 10000000)
	cups = positions[1]
	return cups.Next().Value.(int) * cups.Move(2).Value.(int)
}

func Day23() {

	fmt.Println("Test")
	fmt.Println("Part 1: ", cups1("inputs/day23_test.txt"))
	fmt.Println("Part 2: ", cups2("inputs/day23_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", cups1("inputs/day23.txt"))
	fmt.Println("Part 2: ", cups2("inputs/day23.txt"))
}
