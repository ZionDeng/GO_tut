package main

import "fmt"

func main() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	for country := range countryCapitalMap {
		fmt.Printf("the capital of %s is %s. \n", country, countryCapitalMap[country])
	}
}
