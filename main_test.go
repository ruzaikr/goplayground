package main

import (
	"testing"
)

var testCases = []struct {
	name           string
	inputSlice     [][]int
	expectedOutput int
}{
	{
		"Case 1",
		[][]int{{0, 30}, {5, 10}, {15, 20}},
		2,
	},
	{
		"Case 2",
		[][]int{{7, 10}, {2, 4}},
		1,
	},
	{
		"Case 3",
		[][]int{},
		0,
	},
	{
		"Case 4",
		[][]int{{9, 10}, {4, 9}, {4, 17}},
		2,
	},
	{
		"Case 5",
		[][]int{{1, 5}, {8, 9}, {8, 9}},
		2,
	},
}

func TestMinMeetingRooms(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var actualOutput = minMeetingRooms(testCase.inputSlice)
			if actualOutput != testCase.expectedOutput {
				t.Errorf("got %v, want %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}
