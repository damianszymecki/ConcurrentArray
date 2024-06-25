package main

import (
	"fmt"
	"sync"
)

// Function to print the display
func printDisplay(display [][]int) {
	for _, row := range display {
		for _, pixel := range row {
			if pixel == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Function to add an ASCII character to the display at a given position using concurrency
func addCharToDisplay(display [][]int, char [][]int, x, y int, wg *sync.WaitGroup) {
	defer wg.Done()

	var charWg sync.WaitGroup
	for i := 0; i < len(char); i++ {
		for j := 0; j < len(char[0]); j++ {
			if x+i < len(display) && y+j < len(display[0]) {
				charWg.Add(1)
				go func(i, j int) {
					defer charWg.Done()
					display[x+i][y+j] = char[i][j]
				}(i, j)
			}
		}
	}
	charWg.Wait()
}

func main() {
	display := make([][]int, 64)
	for i := range display {
		display[i] = make([]int, 32)
	}

	charA := [][]int{
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	var wg sync.WaitGroup
	
	wg.Add(1)
	go addCharToDisplay(display, charA, 10, 5, &wg)
	wg.Wait()

	printDisplay(display)
}