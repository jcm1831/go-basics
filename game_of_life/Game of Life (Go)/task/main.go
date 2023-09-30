package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateFirstGeneration(size int) [][]byte {
	generation := make([][]byte, size)
	for i := 0; i < size; i++ {
		generation[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			if populate := rand.Intn(2); populate == 1 {
				generation[i][j] = 'O'
			} else {
				generation[i][j] = ' '
			}
		}
	}
	return generation
}

func countAlive(universe [][]byte) int {
	countAlive := 0
	for i := range universe {
		for j := range universe {
			if universe[i][j] == 'O' {
				countAlive++
			}
		}
	}
	return countAlive
}

func getNumAlive(i, j, size int, universe [][]byte) int {
	neighborIndices := [3]int{-1, 0, 1}
	numAlive := 0
	for _, x := range neighborIndices {
		for _, y := range neighborIndices {
			if x != 0 || y != 0 { // skip origin
				xx := applyPeriodicBoundaries(i+x, size)
				yy := applyPeriodicBoundaries(j+y, size)
				if universe[xx][yy] == 'O' {
					numAlive++
				}
			}
		}
	}
	return numAlive
}

func applyPeriodicBoundaries(val, size int) int {
	if val < 0 {
		return size - 1
	} else if val == size {
		return 0
	} else {
		return val
	}
}

func printUniverse(universe [][]byte, numGeneration int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Generation #%d\n", numGeneration)
	fmt.Printf("Alive: %d\n", countAlive(universe))
	for i := range universe {
		fmt.Println(string(universe[i]))
	}
	time.Sleep(500 * time.Millisecond)
}

func cellSurvives(status byte, numAlive int) bool {
	return status == 'O' && (numAlive == 2 || numAlive == 3)
}

func cellIsReborn(status byte, numAlive int) bool {
	return status == ' ' && numAlive == 3
}

func main() {
	// parse cli input
	var size int
	n, err := fmt.Scan(&size)
	if n != 1 {
		fmt.Println(err)
	}

	// generate current and next generation
	current := generateFirstGeneration(size)
	next := make([][]byte, size)
	for i := range next {
		next[i] = make([]byte, size)
	}
	printUniverse(current, 1)

	// evolve current and next generation
	m := 1
	numGenerations := 20
	for m <= numGenerations {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				numAlive := getNumAlive(i, j, size, current)
				if cellSurvives(current[i][j], numAlive) ||
					cellIsReborn(current[i][j], numAlive) {
					next[i][j] = 'O'
				} else {
					next[i][j] = ' '
				}
			}
		}
		// make deep copy
		for i := range current {
			copy(current[i], next[i])
		}
		m++
		printUniverse(next, m)
	}
}
