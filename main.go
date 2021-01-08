package main

import (
	"fmt"
	"log"
)

func helper(r int, c int, matrix [][]int, visited1s map[string]bool, sizeOfRiver int) (map[string]bool, int) {
	log.Println("helper r,c-> ", r, c)

	if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix[0]){
		return visited1s, sizeOfRiver
	}

	var val = matrix[r][c]

	if val == 0 || visited1s[fmt.Sprintf("%d,%d",r,c)]{
		return visited1s, sizeOfRiver
	}

	sizeOfRiver++ // val must be 1 if not 0
	visited1s[fmt.Sprintf("%d,%d",r,c)] = true

	var upr = r + 1
	var downr = r - 1
	var rightc = c + 1
	var leftc = c - 1

	visited1s, sizeOfRiver = helper(upr, c, matrix, visited1s, sizeOfRiver)
	visited1s, sizeOfRiver = helper(downr, c, matrix, visited1s, sizeOfRiver)
	visited1s, sizeOfRiver = helper(r, rightc, matrix, visited1s, sizeOfRiver)
	visited1s, sizeOfRiver = helper(r, leftc, matrix, visited1s, sizeOfRiver)

	return visited1s, sizeOfRiver

}

func RiverSizes(matrix [][]int) []int {
	var noRows = len(matrix)
	if noRows < 1 {
		return []int{}
	}
	var noCols = len(matrix[0])

	var visited1s = make(map[string]bool)
	var riverSizes = make([]int, 0)

	for r := 0; r < noRows; r++ {
		for c := 0; c < noCols; c++ {
			log.Println("iteration-------------")

			if matrix[r][c] == 0 || visited1s[fmt.Sprintf("%d,%d",r,c)] {
				continue
			}

			// current element is a 1 that has not been processed
			var sizeOfRiver int
			visited1s, sizeOfRiver = helper(r, c, matrix, visited1s, sizeOfRiver)

			riverSizes = append(riverSizes, sizeOfRiver)

		}
	}
	return riverSizes
}

func main() {
	log.Println("hi")

	var matrix = [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	log.Println(RiverSizes(matrix))
}
