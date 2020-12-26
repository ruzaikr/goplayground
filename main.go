package main

import (
	"fmt"
	"sync"
)

type HitCounter struct {
	mu   sync.Mutex
	hits map[int]int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		hits: make(map[int]int),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.mu.Lock()
	this.hits[timestamp] = this.hits[timestamp] + 1
	this.mu.Unlock()
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	this.mu.Lock()
	defer this.mu.Unlock()
	var start = timestamp - 299
	if start < 0 {
		start = 1
	}
	var noOfHits int
	for i := start; i <= timestamp; i++ {
		noOfHits = noOfHits + this.hits[i]
	}
	return noOfHits
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */

func main() {
	var myHitCounter = Constructor()
	myHitCounter.Hit(1)
	myHitCounter.Hit(2)
	myHitCounter.Hit(3)
	fmt.Println(myHitCounter.GetHits(4))

	for i := 0; i < 999; i++ {
		go myHitCounter.Hit(300)
	}

	fmt.Println(myHitCounter.GetHits(300))
	fmt.Println(myHitCounter.GetHits(301))
}
