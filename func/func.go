package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 200
	var ret int

	ret = max(a, b)
	fmt.Println("the max value is: ", ret)
}

func max(a, b int) int {
	// define local variables
	var res int
	if a < b {
		res = b
	} else {
		res = a
	}
	return res
}
