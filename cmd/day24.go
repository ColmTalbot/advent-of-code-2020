package advent

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"strings"
)

func neighbourShifts() map[string]image.Point {
	moves := make(map[string]image.Point)
	moves["h"] = image.Pt(-1, 0)
	moves["u"] = image.Pt(0, -1)
	moves["i"] = image.Pt(1, -1)
	moves["k"] = image.Pt(1, 0)
	moves["m"] = image.Pt(0, 1)
	moves["n"] = image.Pt(-1, 1)
	return moves
}

func findNeighbours(input image.Point) (output []image.Point) {
	for _, shift := range neighbourShifts() {
		output = append(output, input.Add(shift))
	}
	return
}

func loadDirections(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func advanceAlongPath(start image.Point, path string) image.Point {
	moves := neighbourShifts()
	output := image.Pt(start.X, start.Y)
	for character, move := range moves {
		output = output.Add(move.Mul(strings.Count(path, character)))
	}
	return output
}

func translatePath(path string) string {
	path = strings.ReplaceAll(path, "nw", "u")
	path = strings.ReplaceAll(path, "ne", "i")
	path = strings.ReplaceAll(path, "se", "m")
	path = strings.ReplaceAll(path, "sw", "n")
	path = strings.ReplaceAll(path, "w", "h")
	path = strings.ReplaceAll(path, "e", "k")
	return path
}

func countVisited(filename string) map[image.Point]bool {
	paths := loadDirections(filename)
	colours := make(map[image.Point]bool)
	for _, path := range paths {
		end := advanceAlongPath(image.Pt(0, 0), translatePath(path))
		if colours[end] {
			delete(colours, end)
		} else {
			colours[end] = true
		}
	}
	return colours
}

func countBlack(filename string) int {
	colours := countVisited(filename)
	output := len(colours)
	return output
}

func hexagonalLife(filename string) int {
	colours := countVisited(filename)
	nIterations := 100
	for ii := 0; ii < nIterations; ii++ {
		neighbours := make(map[image.Point]int)
		for position := range colours {
			for _, neighbour := range findNeighbours(position) {
				neighbours[neighbour] += 1
			}
		}
		newColours := make(map[image.Point]bool)
		for position := range colours {
			if neighbours[position] > 0 && neighbours[position] < 3 {
				newColours[position] = true
			}
		}
		for position, nn := range neighbours {
			if !colours[position] && nn == 2 {
				newColours[position] = true
			}
		}
		colours = newColours
	}
	return len(colours)
}

func Day24() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", countBlack("inputs/day24_test.txt"))
	fmt.Println("Part 2: ", hexagonalLife("inputs/day24_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", countBlack("inputs/day24.txt"))
	fmt.Println("Part 2: ", hexagonalLife("inputs/day24.txt"))
}
