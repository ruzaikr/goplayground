package main

import (
	"fmt"
)

type Coords struct {
	Row      int
	Col      int
	Distance int
}

func areCoordsValid(coords Coords, maxRow int, maxCol int) bool {
	if coords.Row < 0 || coords.Col < 0 {
		return false
	}
	if coords.Row > maxRow || coords.Col > maxCol {
		return false
	}
	return true
}

/*
Question:
Find the nearest hotels within a range [min, max] in a 2D  grid, where the cells with the value 1 represent part of a
road, a value bigger than 1 is the price of a room per night and an empty cell no road or hotel.
*/
func getNearestHotelsInRange(grid [][]int, startingCoords Coords, minPrice int, maxPrice int) []int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return []int{}
	}

	var nearestHotelsInRange = make([]int, 0)
	var enqueued = make(map[string]bool)
	var bfsQueue = []Coords{startingCoords}

	var currDistance = 0
	for len(bfsQueue) > 0 {

		// dequeue
		var currCoords = bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if currDistance < currCoords.Distance { // we have finished exploring the next current distance level
			currDistance = currCoords.Distance
			if len(nearestHotelsInRange) > 0 {
				return nearestHotelsInRange
			}
		}

		var cellValue = grid[currCoords.Row][currCoords.Col]

		if cellValue > 1 && cellValue >= minPrice && cellValue <= maxPrice {
			nearestHotelsInRange = append(nearestHotelsInRange, grid[currCoords.Row][currCoords.Col])
		}

		if cellValue == 0 {
			continue
		}

		var newDistance = currCoords.Distance + 1
		var northCoords = Coords{Row: currCoords.Row - 1, Col: currCoords.Col, Distance: newDistance}
		var eastCoords = Coords{Row: currCoords.Row, Col: currCoords.Col + 1, Distance: newDistance}
		var southCoords = Coords{Row: currCoords.Row + 1, Col: currCoords.Col, Distance: newDistance}
		var westCoords = Coords{Row: currCoords.Row, Col: currCoords.Col - 1, Distance: newDistance}
		var neighbors = []Coords{northCoords, eastCoords, southCoords, westCoords}

		for i := 0; i < len(neighbors); i++ {
			var key = fmt.Sprintf("%d,%d", neighbors[i].Col, neighbors[i].Row)
			if areCoordsValid(neighbors[i], len(grid)-1, len(grid[0])-1) && !enqueued[key] {
				bfsQueue = append(bfsQueue, neighbors[i]) // enqueue
				enqueued[key] = true
			}
		}
	}

	return nearestHotelsInRange
}

func main() {
	var startingCol = 3
	var startingRow = 2

	var grid = [][]int{
		{6, 4, 0, 15, 7},
		{2, 4, 1, 0, 1},
		{6, 1, 1, 1, 1},
		{9, 1, 1, 1, 0},
		{0, 1, 0, 1, 0},
	}

	//  var grid = [][]int{
	//     {2, 8, 0},
	//     {9, 1, 1},
	//     {3, 1, 1},
	// }

	fmt.Println(getNearestHotelsInRange(grid, Coords{Row: startingRow, Col: startingCol, Distance: 0}, 12, 15))
}
