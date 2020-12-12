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
		for j := i + 1; j < len(websites); j++ {
			if websites[j] == websites[i] {
				continue
			}
			for k := j + 1; k < len(websites); k++ {
				if websites[k] == websites[j] {
					continue
				}
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
		for i := 0; i < len(websites); i++ {
			for j := i + 1; j < len(websites); j++ {
				if websites[j] == websites[i] {
					continue
				}
				for k := j + 1; k < len(websites); k++ {
					if websites[k] == websites[j] {
						continue
					}
					var seq = fmt.Sprintf("%s,%s,%s",websites[i],websites[j],websites[k])
					var count, _ = threeSeqs[seq]
					var newCount = count + 1
					threeSeqs[seq] = newCount
					if newCount > highestCount {
						highestCount = newCount
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
