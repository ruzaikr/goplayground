package main

import (
	"log"
)

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		log.Fatalln("Length of input array must be larger than or equal to 2")
	}

	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			var sum = nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {

}
