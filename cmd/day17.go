package advent

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func emptyBoard(width, height int) (board [][][]bool) {
	for ii := 0; ii < height; ii++ {
		var temp [][]bool
		for jj := 0; jj < width; jj++ {
			temp = append(temp, make([]bool, width))
		}
		board = append(board, temp)
	}
	return
}

func countAliveNeighbours(board [][][]bool, ii, jj, kk int) (neighbours int) {
	for _ii := ii - 1; _ii <= ii+1; _ii++ {
		if _ii < 0 || _ii >= len(board) {
			continue
		}
		for _jj := jj - 1; _jj <= jj+1; _jj++ {
			if _jj < 0 || _jj >= len(board[0]) {
				continue
			}
			for _kk := kk - 1; _kk <= kk+1; _kk++ {
				if _kk < 0 || _kk >= len(board[0][0]) {
					continue
				} else if _ii == ii && _jj == jj && _kk == kk {
					continue
				} else if board[_ii][_jj][_kk] {
					neighbours += 1
				}
			}
		}
	}
	return
}

func countAlive(board [][][]bool) (alive [][][]int) {
	alive = make([][][]int, len(board))
	for ii, level := range board {
		alive[ii] = make([][]int, len(level))
		for jj, line := range level {
			alive[ii][jj] = make([]int, len(line))
			for kk := range line {
				alive[ii][jj][kk] = countAliveNeighbours(board, ii, jj, kk)
			}
		}
	}
	return
}

func countActive(board [][][]bool) (active int) {
	for ii, level := range board {
		for jj, line := range level {
			for kk := range line {
				if board[ii][jj][kk] {
					active++
				}
			}
		}
	}
	return
}

func updateBoard(board [][][]bool, alive [][][]int) [][][]bool {
	var change bool
	for ii, level := range board {
		for jj, line := range level {
			for kk := range line {
				if board[ii][jj][kk] && !(alive[ii][jj][kk] == 2 || alive[ii][jj][kk] == 3) {
					change = true
				} else if !board[ii][jj][kk] && alive[ii][jj][kk] == 3 {
					change = true
				} else {
					change = false
				}
				if change {
					board[ii][jj][kk] = !board[ii][jj][kk]
				}
			}
		}
	}
	return board
}

func boardFromFile(filename string) (board [][][]bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	initialState := strings.Split(string(data), "\n")

	width := len(initialState) + 6*2
	height := 1 + 6*2

	board = emptyBoard(width, height)

	for ii, line := range initialState {
		for jj, character := range line {
			if character == '#' {
				board[6][6+ii][6+jj] = true
			}
		}
	}
	return
}

func playGame(filename string) (active int) {
	board := boardFromFile(filename)
	for ii := 0; ii < 6; ii++ {
		neighbours := countAlive(board)
		board = updateBoard(board, neighbours)
		active = countActive(board)
	}
	return
}

func emptyBoard2(width, height int) (board [][][][]bool) {
	for zz := 0; zz < height; zz++ {
		board = append(board, emptyBoard(width, height))
	}
	return
}

func countAliveNeighbours2(board [][][][]bool, zz, ii, jj, kk int) (neighbours int) {
	for _zz := zz - 1; _zz <= zz+1; _zz++ {
		if _zz < 0 || _zz >= len(board) {
			continue
		}
		for _ii := ii - 1; _ii <= ii+1; _ii++ {
			if _ii < 0 || _ii >= len(board[0]) {
				continue
			}
			for _jj := jj - 1; _jj <= jj+1; _jj++ {
				if _jj < 0 || _jj >= len(board[0][0]) {
					continue
				}
				for _kk := kk - 1; _kk <= kk+1; _kk++ {
					if _kk < 0 || _kk >= len(board[0][0][0]) {
						continue
					} else if _ii == ii && _jj == jj && _kk == kk && _zz == zz {
						continue
					} else if board[_zz][_ii][_jj][_kk] {
						neighbours += 1
					}
				}
			}
		}
	}
	return
}

func countAlive2(board [][][][]bool) (alive [][][][]int) {
	alive = make([][][][]int, len(board))
	for zz, space := range board {
		alive[zz] = make([][][]int, len(space))
		for ii, level := range space {
			alive[zz][ii] = make([][]int, len(level))
			for jj, line := range level {
				alive[zz][ii][jj] = make([]int, len(line))
				for kk := range line {
					alive[zz][ii][jj][kk] = countAliveNeighbours2(board, zz, ii, jj, kk)
				}
			}
		}
	}
	return
}

func countActive2(board [][][][]bool) (active int) {
	for _, space := range board {
		active += countActive(space)
	}
	return
}

func updateBoard2(board [][][][]bool, alive [][][][]int) [][][][]bool {
	for zz, space := range board {
		updateBoard(space, alive[zz])
	}
	return board
}

func boardFromFile2(filename string) (board [][][][]bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	initialState := strings.Split(string(data), "\n")

	expand := 6
	width := len(initialState) + expand*2
	height := 1 + expand*2

	board = emptyBoard2(width, height)

	for ii, line := range initialState {
		for jj, character := range line {
			if character == '#' {
				board[expand][expand][expand+ii][expand+jj] = true
			}
		}
	}
	return
}

func playGame2(filename string) (active int) {
	board := boardFromFile2(filename)
	for ii := 0; ii < 6; ii++ {
		neighbours := countAlive2(board)
		board = updateBoard2(board, neighbours)
		active = countActive2(board)
	}
	return
}

func Day17() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", playGame("inputs/day17_test.txt"))
	fmt.Println("Part 2: ", playGame2("inputs/day17_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", playGame("inputs/day17.txt"))
	fmt.Println("Part 2: ", playGame2("inputs/day17.txt"))
}
