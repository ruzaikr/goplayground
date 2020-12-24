package main

import (
	"fmt"
)

// Coloring A Border
// =================
//

func helper(grid [][]int, r int, c int, oldColor int, newColor int, visited map[string]bool) [][]int {
	var key = fmt.Sprintf("%d,%d", r, c)
	if _, exists := visited[key]; exists {
		return grid
	}
	visited[key] = true // mark as visited

	if r == -1 || r == len(grid) || c == -1 || c == len(grid[r]) { // check if r,c is out of bounds
		return grid
	}

	if grid[r][c] != oldColor {
		return grid
	}

	var leftc = c - 1
	var rightc = c + 1
	var upr = r - 1
	var downr = r + 1

	if r == 0 || r == (len(grid)-1) || c == 0 || c == (len(grid[r])-1) ||
		grid[r][leftc] != oldColor ||
		grid[r][rightc] != oldColor ||
		grid[upr][c] != oldColor ||
		grid[downr][c] != oldColor {
		grid[r][c] = newColor
	}

	helper(grid, r, leftc, oldColor, newColor, visited)
	helper(grid, r, rightc, oldColor, newColor, visited)
	helper(grid, upr, c, oldColor, newColor, visited)
	helper(grid, downr, c, oldColor, newColor, visited)

	return grid
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {

	// @todo: case where r0,c0 does not have any adjacent squares with the same color

	return helper(grid, r0, c0, grid[r0][c0], color, make(map[string]bool))
}

func main() {

}
