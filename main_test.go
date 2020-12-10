package main

import (
	"testing"
)

var twoSumTests = []struct {
	name string
	inputArray  []int
	target int
	expectedOutput []int
}{
	{
	"Case 1",
	[]int{2, 7, 11, 15},
		9,
	[]int{0, 1},
	},
	{
	"Case 2",
	[]int{3, 2, 4},
		6,
	[]int{1, 2},
	},
}

func areIntArraysEqual(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}


func TestTwoSum(t *testing.T) {
	for _, twoSumTest := range twoSumTests {
		t.Run(twoSumTest.name, func(t *testing.T) {
			var actualOutput = twoSum(twoSumTest.inputArray, twoSumTest.target)
			if !areIntArraysEqual(actualOutput, twoSumTest.expectedOutput) {
				t.Errorf("got %v, want %v", actualOutput, twoSumTest.expectedOutput)
			}
		})
	}
}