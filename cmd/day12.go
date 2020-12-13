package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func advanceAlongDirection(xposition int, yposition int, operand uint8, step int) (int, int) {
	if operand == 'N' {
		yposition += step
	} else if operand == 'S' {
		yposition -= step
	} else if operand == 'E' {
		xposition += step
	} else if operand == 'W' {
		xposition -= step
	}
	return xposition, yposition
}

func fullManhattanDistance1(filename string) int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	var xposition, yposition, direction int
	var step int64
	var operand uint8
	directions := []uint8{'E', 'S', 'W', 'N'}

	for _, instruction := range instructions {
		operand = instruction[0]
		step, err = strconv.ParseInt(instruction[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if operand == 'N' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, operand, int(step))
		} else if operand == 'S' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, operand, int(step))
		} else if operand == 'E' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, operand, int(step))
		} else if operand == 'W' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, operand, int(step))
		} else if operand == 'L' {
			direction = (direction - int(step)/90 + 4) % 4
		} else if operand == 'R' {
			direction = (direction + int(step)/90) % 4
		} else if operand == 'F' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, directions[direction], int(step))
		}
	}
	return AbsInt(xposition) + AbsInt(yposition)
}

func fullManhattanDistance2(filename string) int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	var xposition, yposition int
	xwaypoint := 10
	ywaypoint := 1
	var step int64
	var operand uint8

	for _, instruction := range instructions {
		operand = instruction[0]
		step, err = strconv.ParseInt(instruction[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if operand == 'N' {
			xwaypoint, ywaypoint = advanceAlongDirection(xwaypoint, ywaypoint, operand, int(step))
		} else if operand == 'S' {
			xwaypoint, ywaypoint = advanceAlongDirection(xwaypoint, ywaypoint, operand, int(step))
		} else if operand == 'E' {
			xwaypoint, ywaypoint = advanceAlongDirection(xwaypoint, ywaypoint, operand, int(step))
		} else if operand == 'W' {
			xwaypoint, ywaypoint = advanceAlongDirection(xwaypoint, ywaypoint, operand, int(step))
		} else if operand == 'L' {
			switch int(step) {
			case 90:
				xwaypoint, ywaypoint = -ywaypoint, xwaypoint
			case 180:
				xwaypoint, ywaypoint = -xwaypoint, -ywaypoint
			case 270:
				xwaypoint, ywaypoint = ywaypoint, -xwaypoint
			}
		} else if operand == 'R' {
			switch int(step) {
			case 90:
				xwaypoint, ywaypoint = ywaypoint, -xwaypoint
			case 180:
				xwaypoint, ywaypoint = -xwaypoint, -ywaypoint
			case 270:
				xwaypoint, ywaypoint = -ywaypoint, xwaypoint
			}
		} else if operand == 'F' {
			xposition, yposition = advanceAlongDirection(xposition, yposition, 'E', xwaypoint*int(step))
			xposition, yposition = advanceAlongDirection(xposition, yposition, 'N', ywaypoint*int(step))
		}
	}
	return AbsInt(xposition) + AbsInt(yposition)
}

func Day12() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", fullManhattanDistance1("inputs/day12_test.txt"))
	fmt.Println("Part 2: ", fullManhattanDistance2("inputs/day12_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", fullManhattanDistance1("inputs/day12.txt"))
	fmt.Println("Part 2: ", fullManhattanDistance2("inputs/day12.txt"))
}
