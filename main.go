package main

import (
	"fmt"
)

// Coloring A Border
// =================
//

func helper(grid [][]int, r int, c int, oldColor int, visited map[string]bool, border [][]int) ([][]int, [][]int) {
	var key = fmt.Sprintf("%d,%d", r, c)
	if _, exists := visited[key]; exists {
		return grid, border
	}
	visited[key] = true // mark as visited

	if r == -1 || r == len(grid) || c == -1 || c == len(grid[r]) { // check if r,c is out of bounds
		return grid, border
	}

	if grid[r][c] != oldColor {
		return grid, border
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
		border = append(border, []int{r, c})
	}

	grid, border = helper(grid, r, leftc, oldColor, visited, border)
	grid, border = helper(grid, r, rightc, oldColor, visited, border)
	grid, border = helper(grid, upr, c, oldColor, visited, border)
	grid, border = helper(grid, downr, c, oldColor, visited, border)

	return grid, border
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {

	// @todo: case where r0,c0 does not have any adjacent squares with the same color

	if len(grid) < 1 || len(grid[0]) < 1 {
		return grid
	}

	var border = make([][]int, 0, len(grid)*len(grid[0]))
	grid, border = helper(grid, r0, c0, grid[r0][c0], make(map[string]bool), border)

	for i := 0; i < len(border); i++ {
		grid[border[i][0]][border[i][1]] = color
	}

	return grid
}

func main() {

}
