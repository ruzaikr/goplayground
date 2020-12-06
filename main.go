package main

import (
	"log"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			intervals = append(intervals[:i], intervals[i+1:]...)
		}
	}
	return intervals
}

func main() {
	var myIntervals = [][]int{{1,3},{2,6},{8,10},{15,18}}
	log.Println(merge(myIntervals))
}
