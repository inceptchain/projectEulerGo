package main

import (
	"fmt"
)

func main() {
	//si := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gridSize := 21
	grid := make([][]int, gridSize)
	// for i := range grid {
	// 	grid[i] = si[:gridSize]
	// }
	//fmt.Println(len(grid))

	for y := 0; y < gridSize; y++ {
		grid[y] = make([]int, gridSize)
		for x := 0; x < gridSize; x++ {

			if x == 0 || y == 0 {
				grid[y][x] = 1

			} else {
				grid[y][x] = grid[y][x-1] + grid[y-1][x]
			}

			if x == y {
				fmt.Println(grid[y][x])
			}
		}

	}

}
