package main

import (
	"testing"
)

var testCases = []struct {
	name           string
	usernames      []string
	timestamps     []int
	websites       []string
	expectedOutput []string
}{
	{
	"Case 1",
	[]string{"joe","joe","joe","james","james","james","james","mary","mary","mary"},
		[]int{1,2,3,4,5,6,7,8,9,10},
	[]string{"home","about","career","home","cart","maps","home","home","about","career"},
	[]string{"home","about","career"},
	},
}

func areStringArraysEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMostVisitedPattern(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var actualOutput = mostVisitedPattern(testCase.usernames, testCase.timestamps, testCase.websites)
			if !areStringArraysEqual(actualOutput, testCase.expectedOutput) {
				t.Errorf("got %v, want %v", actualOutput, testCase.expectedOutput)
			}
		})
	}
}