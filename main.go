package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		var temp = a
		a = b
		b = temp
	}

	var res = ""

	var c int
	for i := 0; i < len(a) || c == 1; i++ {
		var ai = len(a)-1-i
		var bi = len(b)-1-i
		var r = "0"
		var s int

		if i < len(b) {
			var temp, _ = strconv.Atoi(string(b[bi]))
			s += temp
		}

		if i < len(a) {
			var temp, _ = strconv.Atoi(string(a[ai]))
			s += temp
		}

		s += c

		if s == 3 {
			c = 1
			r = "1"
		}else if s == 2 {
			c = 1
		}else if s == 1 {
			c = 0
			r = "1"
		}else {
			c = 0
		}

		res = r + res
	}
	return res
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
