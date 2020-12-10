package main

import (
	"testing"
)

var mergeTests = []struct {
	name string
	inputParams  [][]int
	expectedOutput [][]int
}{
	{
		"Case 1",
		[][]int{{1,3},{2,6},{8,10},{15,18}},
		[][]int{{1,6},{8,10},{15,18}}},
	{
		"Case 2",
		[][]int{{1,4},{4,5}},
		[][]int{{1,5}},
	},
	{
		"Case 3",
		[][]int{{1,4},{0,4}},
		[][]int{{0,4}},
	},
	{
		"Case 4", // required sorting
		[][]int{{1,4},{0,1}},
		[][]int{{0,4}},
	},
	{
		"Case 5",
		[][]int{{1,4},{2,3}},
		[][]int{{1,4}},
	},
	{
		"Case 6",
		[][]int{{1,4},{0,2},{3,5}},
		[][]int{{0,5}},
	},
	{
		"Case 7",
		[][]int{{2,3},{4,5},{6,7},{8,9},{1,10}},
		[][]int{{1,10}},
	},
	{
		"Case 8",
		[][]int{{1,3},{0,2},{2,3},{4,6},{5,5},{0,2},{3,3}},
		[][]int{{0,3},{4,6}},
	},
}

func areIntervalsEqual(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	for _, mergeTest := range mergeTests {
		t.Run(mergeTest.name, func(t *testing.T) {
			var actualOutput = Merge(mergeTest.inputParams)
			if !areIntervalsEqual(actualOutput, mergeTest.expectedOutput) {
				t.Errorf("got %v, want %v", actualOutput, mergeTest.expectedOutput)
			}
		})
	}
}