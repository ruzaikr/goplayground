package main

import (
	"fmt"
	"sort"
)

type visit struct {
	timestamp int
	username string
	website string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	var websitesMap = make(map[string]bool)
	var websites = make([]string, 0)
	var visits = make([]visit, len(timestamp))
	for i := 0; i < len(timestamp); i++ {
		visits[i] = visit{
			username: username[i],
			timestamp: timestamp[i],
			website: website[i],
		}

		if _, exists := websitesMap[website[i]]; !exists {
			websitesMap[website[i]] = true
			websites = append(websites, website[i])
		}
	}

	var threeSeqs = make(map[string]int)

	// make all possible 3-sequences
	for i := 0; i < len(websites); i++ {
		for j := 0; j < len(websites); j++ {
			for k := 0; k < len(websites); k++ {
				var seq = fmt.Sprintf("%s,%s,%s",websites[i],websites[j],websites[k])
				threeSeqs[seq] = 0
			}
		}
	}

	sort.Slice(visits, func(i, j int) bool {
		return visits[i].timestamp < visits[j].timestamp
	})

	var usernameToVisits = make(map[string][]string)

	for i := 0; i < len(visits); i++ {
		var v, exists = usernameToVisits[visits[i].username]
		if exists {
			usernameToVisits[visits[i].username] = append(v, visits[i].website)
		}else {
			usernameToVisits[visits[i].username] =  []string{visits[i].website}
		}
	}

	var highestCount = 0
	var mostVisitedPattenSlice = make([]string, 3)
	for _, websites := range usernameToVisits {
		var threeSeqsUser = make(map[string]bool)
		for i := 0; i < len(websites); i++ {
			for j := i + 1; j < len(websites); j++ {
				for k := j + 1; k < len(websites); k++ {
					var seq = fmt.Sprintf("%s,%s,%s",websites[i],websites[j],websites[k])
					if _, exists := threeSeqsUser[seq]; !exists {
						threeSeqsUser[seq] = true
					}else {
						continue
					}
					var count, _ = threeSeqs[seq]
					var newCount = count + 1
					threeSeqs[seq] = newCount
					if newCount > highestCount {
						highestCount = newCount
						mostVisitedPattenSlice = []string{websites[i], websites[j], websites[k]}
					}else if newCount == highestCount && seq < fmt.Sprintf("%s,%s,%s",mostVisitedPattenSlice[0],mostVisitedPattenSlice[1],mostVisitedPattenSlice[2]) {
						mostVisitedPattenSlice = []string{websites[i], websites[j], websites[k]}
					}
				}
			}
		}
	}

	return mostVisitedPattenSlice
}

func main() {

}
