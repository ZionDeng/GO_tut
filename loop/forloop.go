package main

import "fmt"

func main() {
	strings := []string{"go", "rn"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	num := [6]int{1, 2, 3,5}
	for i, x := range num {
		fmt.Printf("the %d th x is : %d. ", i, x)

	}
}
