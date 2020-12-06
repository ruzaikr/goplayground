package main

import (
	"log"
	"sort"
)

type byCustom [][]int

func (intervals byCustom) Len() int {
	return len(intervals)
}
func (intervals byCustom) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}
func (intervals byCustom) Less(i, j int) bool {
	return intervals[i][0] < intervals[j][0]
}

func removeFromIntervals(i int, intervals [][]int) [][]int {
	return append(intervals[:i], intervals[i+1:]...)
}

func isAContainedWithinB(a []int, b []int) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Sort(byCustom(intervals))

	for i := 0; i < len(intervals) - 1; i++ {
		if isAContainedWithinB(intervals[i], intervals[i+1]) {
			intervals = removeFromIntervals(i, intervals)
		}else if isAContainedWithinB(intervals[i+1], intervals[i]) {
			intervals = removeFromIntervals(i+1, intervals)
		} else if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			intervals = removeFromIntervals(i, intervals)
		}
	}
	return intervals
}

func main() {
	var testCases = [][][]int{
		{{1,3},{2,6},{8,10},{15,18}},
		{{1,4},{4,5}},
	}

	for i := 0; i < len(testCases); i++ {
		log.Println(Merge(testCases[i]))
	}
}
