package main

import (
	"log"
)

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		log.Fatalln("Length of input array must be larger than or equal to 2")
	}

	var m = make(map[int]int) // key: complement, val: index

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		var c = target - nums[i]
		var v, exists = m[c]
		if exists && v != i {
			return []int{i,v}
		}
	}

	return []int{}
}

func main() {

}
