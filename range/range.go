package main

import "fmt"

func main() {
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cat"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)

	}

}
