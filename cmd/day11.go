package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countNeighbours(seatMap [][]rune, ii int, jj int) (neighbours int) {
	line := seatMap[jj]
	for mm := ii - 1; mm <= ii+1; mm++ {
		if mm < 0 || mm >= len(line) {
			continue
		}
		for nn := jj - 1; nn <= jj+1; nn++ {
			if nn < 0 || nn >= len(seatMap) || (mm == ii && nn == jj) {
				continue
			}
			if seatMap[nn][mm] == '#' {
				neighbours++
			}
		}
	}
	return
}

func lookAlongDirection(seatMap [][]rune, ii int, jj int) (found int) {
	dx := []int{1, 1, 0, -1, -1, -1, 0, 1}
	dy := []int{0, -1, -1, -1, 0, 1, 1, 1}
	var xpos, ypos int
	for idx := range dx {
		xpos = jj + dx[idx]
		ypos = ii + dy[idx]
		for xpos >= 0 && ypos >= 0 && ypos < len(seatMap[ii]) && xpos < len(seatMap) {
			if seatMap[xpos][ypos] == '#' {
				found++
				break
			} else if seatMap[xpos][ypos] == 'L' {
				break
			}
			xpos += dx[idx]
			ypos += dy[idx]
		}
	}
	return
}

func countAllNeighbours(seatMap [][]rune, neighbourFunc func([][]rune, int, int) int) (seats [][]int) {
	for ii, line := range seatMap {
		seats = append(seats, make([]int, 0))
		for jj := range line {
			seats[ii] = append(seats[ii], neighbourFunc(seatMap, jj, ii))
		}
	}
	return
}

func updateSeats(seats [][]rune, neighbours [][]int, maximum int) (newSeats [][]rune) {
	var newChar rune
	for ii := range seats {
		newSeats = append(newSeats, make([]rune, 0, len(seats[ii])))
		for jj := range seats[ii] {
			if seats[ii][jj] == 'L' && neighbours[ii][jj] == 0 {
				newChar = '#'
			} else if seats[ii][jj] == '#' && neighbours[ii][jj] >= maximum {
				newChar = 'L'
			} else {
				newChar = seats[ii][jj]
			}
			newSeats[ii] = append(newSeats[ii], newChar)
		}
	}
	return
}

func countOccupied(seats [][]rune) (occupied int) {
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				occupied++
			}
		}
	}
	return
}

func printSeats(seats [][]rune) {
	for _, row := range seats {
		for _, seat := range row {
			fmt.Print(string(seat))
		}
		fmt.Print("\n")
	}
}

func printNeighbours(seats [][]int) {
	for _, row := range seats {
		for _, seat := range row {
			fmt.Print(seat)
		}
		fmt.Print("\n")
	}
}

func findStableOccupation(filename string, neighbourFunc func([][]rune, int, int) int, maximum int) (occupied int) {
	allData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	seatMap := strings.Split(string(allData), "\n")
	var seats [][]rune
	var neighbourMap [][]int
	for ii := range seatMap {
		neighbourMap = append(neighbourMap, make([]int, 0, len(seatMap[ii])))
		seats = append(seats, make([]rune, 0, len(seatMap[ii])))
		for jj := range seatMap[ii] {
			neighbourMap[ii] = append(neighbourMap[ii], 0)
			seats[ii] = append(seats[ii], rune(seatMap[ii][jj]))
		}
	}
	oldOccupied := countOccupied(seats)
	occupied = -1
	for oldOccupied != occupied {
		oldOccupied = occupied
		neighbourMap = countAllNeighbours(seats, neighbourFunc)
		seats = updateSeats(seats, neighbourMap, maximum)
		occupied = countOccupied(seats)
	}
	return
}

func Day11() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", findStableOccupation("inputs/day11_test.txt", countNeighbours, 4))
	fmt.Println("Part 2: ", findStableOccupation("inputs/day11_test.txt", lookAlongDirection, 5))
	fmt.Println("Main")
	fmt.Println("Part 1: ", findStableOccupation("inputs/day11.txt", countNeighbours, 4))
	fmt.Println("Part 2: ", findStableOccupation("inputs/day11.txt", lookAlongDirection, 5))
}
