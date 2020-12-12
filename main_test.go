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
	{
	"Case 2",
	[]string{"u1","u1","u1","u2","u2","u2"},
		[]int{1,2,3,4,5,6},
	[]string{"a","b","c","a","b","a"},
	[]string{"a","b","a"},
	},
	{
		"Case 3",
		[]string{"dowg","dowg","dowg"},
		[]int{158931262,562600350,148438945},
		[]string{"y","loedo","y"},
		[]string{"y","y","loedo"},
	},
	{
		"Case 4",
		[]string{"h","eiy","cq","h","cq","txldsscx","cq","txldsscx","h","cq","cq"},
		[]int{527896567,334462937,517687281,134127993,859112386,159548699,51100299,444082139,926837079,317455832,411747930},
		[]string{"hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","hibympufi","yljmntrclw","hibympufi","yljmntrclw"},
		[]string{"hibympufi","hibympufi","yljmntrclw"},
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