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

func isSquare(i int, j int, sideLen int, matrix [][]byte) bool {
	for k := i; k <= sideLen; k++ {
		for l := j; l <= sideLen; l++ {
			fmt.Println(matrix[k][l])
			if matrix[k][l] != 49 {
				return false
			}
		}
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

	fmt.Println("maxSquareSide = ", maxSquareSide)

	outermost:
	for maxSquareSide > 0 {
		for y := 0; y < (lenY - maxSquareSide + 1); y++ {
			for x := 0; x < (lenX - maxSquareSide + 1); x++ {
				if isSquare(y, x, maxSquareSide, matrix) {
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
	fmt.Println("Hello World")
}
