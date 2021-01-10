// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputArr = []string{"P1", "D1", "P2", "D2"}  // true
	var inputArr2 = []string{"P1", "P2", "D1", "D2"} // true
	var inputArr3 = []string{"P2", "P1", "D1", "D2"} // true
	var inputArr4 = []string{"P1", "D1", "P2"}       // false
	var inputArr5 = []string{"D1", "P1"}             // false
	var inputArr6 = []string{"P1", "D1", "P1", "D2"} // false
	var inputArr7 = []string{"P1", "D1", "P1", "D1"} // false

	fmt.Println(determineScheduleIsValid(inputArr))
	fmt.Println(determineScheduleIsValid(inputArr2))
	fmt.Println(determineScheduleIsValid(inputArr3))
	fmt.Println(determineScheduleIsValid(inputArr4))
	fmt.Println(determineScheduleIsValid(inputArr5))
	fmt.Println(determineScheduleIsValid(inputArr6))
	fmt.Println(determineScheduleIsValid(inputArr7))
}

/*

A valid delivery is if the dropoff is after the pickup


P1, D1, P2, D2 => valid
P1, P2, D1, D2 => valid
P2, P1, D1, D2 => valid

P1, D1, P1, D1 => valid


P1, D1, P2 => invalid
D1, P1 => invalid
P1, D1, P1, D2 => invalid




Your goal is to determine if a schedule is valid.

input parameter : ["P1", "D1...


*/

func determineScheduleIsValid(inputArr []string) bool {
	var mapPickupsToDeliveries = make(map[string]bool)
	for i := 0; i < len(inputArr); i++ {
		if strings.HasPrefix(inputArr[i], "P") {
			var indexOfPickup = inputArr[i][1:]
			if mapPickupsToDeliveries[indexOfPickup] {
				return false
			}
			mapPickupsToDeliveries[indexOfPickup] = false
		}
		if strings.HasPrefix(inputArr[i], "D") {
			var indexOfDelivery = inputArr[i][1:]
			var val, exists = mapPickupsToDeliveries[indexOfDelivery]
			if exists && val == false {
				delete(mapPickupsToDeliveries, indexOfDelivery)
			} else {
				return false
			}
		}
	}

	return len(mapPickupsToDeliveries) == 0
}

// O(n)
// O(n)

// func determineScheduleIsValid() bool {
//   for i := 0; i < len(inputArr); i++ {
//     if strings.HasPrefix(inputArr[i], "P") {
//       var pickupIsValid bool
//       var indexOfPickup = inputArr[i][1:]
//       for j := i + 1; j < len(inputArr); j++ {
//         if strings.HasPrefix(inputArr[j], "D") {
//           var indexOfDelivery = inputArr[j][1:]
//           if indexOfPickup == indexOfDelivery {
//             pickupIsValid = true
//           }
//         }
//       }
//       if !pickupIsValid {
//         return false
//       }
//     }
//   }
//   return true
// }

// invalid case [P1, D2, P2, D1]  invalid
// hash table { 1: false,
//
// invalid case [P1, D1, P2, D3]  invalid
// hash table { 1: true, 2: false
//

// time complexity: [P1, P2... P8]
// O(n^2)

/*

Given n, generate all valid unique schedules

if n = 1
  P1, D1

if n = 2
  P1, D1, P2, D2
  P1, P2, D1, D2
  P1, P2, D2, D1
  P2, P1, D1, D2
  P2, P1, D2, D1

[P/D 1-n/2 ], [ ], [ ], [ ]
n/2

----------


n = 2
choices = [P1, P2]

    choices= [P2, D1]  (P1)

         choices [D1, D2]  (P1, P2)


            choices [D2]   (P1, P2, D1)

                 choices [] (P1, P2, D1, D2)



*/

func generateAllValidSchedules(n int) {
	var result = make([]string, 0)

}
