package main

import (
	"fmt"
	"strings"
)

func isEditDistanceOne(a, b string) bool {
	// if strings are equal
	if a == b {
		return true
	}

	// lengths are equal but strings are not equal
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				continue
			}
			b = strings.Replace(b, string(b[i]), string(a[i]), 1)
			if a == b {
				return true
			}
			return false
		}
	}

	var larger, smaller string
	if len(b) > len(a) {
		larger = b
		smaller = a
	} else {
		larger = a
		smaller = b
	}

	if (len(larger) - len(smaller)) > 1 {
		return false
	}

	// lengths differ by 1
	if string(larger[:len(larger)-1]) == smaller {
		return true
	}
	for i := 0; i < len(smaller); i++ {
		if larger[i] == smaller[i] {
			continue
		}
		var trimmedLarger = strings.Replace(larger, string(larger[i]), "", 1)
		if trimmedLarger == smaller {
			return true
		}
		return false
	}

	return false
}

func main() {
	fmt.Println(isEditDistanceOne("bexac", "bexac"))
	fmt.Println(isEditDistanceOne("cait", "caita"))
}
