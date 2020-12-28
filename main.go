package main

import (
	"fmt"
)

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSquare(x int, y int, sideLen int, matrix [][]byte) bool {
	fmt.Println("x, y: ", x, y, sideLen)
	if sideLen == 1 {
		return matrix[y][x] == 49
	}
	for i := 0; i < sideLen; i++ {
		var innerY = y
		for j := 0; j < sideLen; j++ {
			fmt.Println("yelo", x, innerY, matrix[x][innerY])
			if matrix[x][innerY] != 49 {
				return false
			}
			innerY = innerY + 1
		}
		x = x + 1
	}
	return true
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) < 1 || len(matrix[0]) < 1{
		return 0
	}

	var lenY = len(matrix)
	var lenX = len(matrix[0])

	var maxSquareSide = getMin(lenX, lenY)


	outermost:
	for maxSquareSide > 0 {
		fmt.Println("maxSquareSide = ", maxSquareSide)
		for y := 0; y < (lenY - maxSquareSide + 1); y++ {
			for x := 0; x < (lenX - maxSquareSide + 1); x++ {
				if isSquare(x, y, maxSquareSide, matrix) {
					fmt.Println("y = ", y, " x = ", x)
					break outermost
				}
			}
		}
		maxSquareSide = maxSquareSide - 1
	}

	return maxSquareSide*maxSquareSide
}


func main() {
	var line1 = []byte("00")
	var input [][]byte
	input = append(input, line1)
	fmt.Println(maximalSquare(input))
}

// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
