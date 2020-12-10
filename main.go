package main

import (
	"log"
	"sort"
)

type visit struct {
	timestamp int
	username string
	website string
}


func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	var visits = make([]visit, len(timestamp))
	for i := 0; i < len(timestamp); i++ {
		visits[i] = visit{
			username: username[i],
			timestamp: timestamp[i],
			website: website[i],
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

	// make 3 sequences
	//var threeSeqs = make(map[string]int)

	for username, visits := range usernameToVisits {
		log.Println("chimken", username, visits)
	}


	return []string{}
}

func main() {

}
