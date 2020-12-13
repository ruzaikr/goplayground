package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	} else if len(intervals) == 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	var intervalsAllocated = make(map[int]bool)

	var noOfMeetingRooms int

	for i := 0; i < len(intervals); i++ {

		var keysToAllocate []int
		if intervalsAllocated[i] {
			continue
		}
		keysToAllocate = append(keysToAllocate, i)
		var allocatedJMap = make(map[string]bool)
		for j := i + 1; j < len(intervals); j++ {
			var keyJ = fmt.Sprintf("%d,%d", intervals[j][0], intervals[j][1])
			if intervalsAllocated[j] || allocatedJMap[keyJ] {
				continue
			}

			if intervals[i][1] <= intervals[j][0] {
				keysToAllocate = append(keysToAllocate, j)
				allocatedJMap[keyJ] = true
			}

		}

		for i := 0; i < len(keysToAllocate); i++ {
			intervalsAllocated[keysToAllocate[i]] = true
		}

		noOfMeetingRooms = noOfMeetingRooms + 1

	}

	return noOfMeetingRooms
}

func main() {

}
