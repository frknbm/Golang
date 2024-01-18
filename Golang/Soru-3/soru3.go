package main

import (
	"fmt"
)

func mostRepeated(data []string) string {
	counts := make(map[string]int)

	for _, item := range data {
		counts[item]++
	}

	maxCount := 0
	var mostRepeatedItem string

	for item, count := range counts {
		if count > maxCount {
			maxCount = count
			mostRepeatedItem = item
		}
	}

	return mostRepeatedItem
}

func main() {
	input := []string{"apple", "pie", "apple", "red", "red", "red"}
	output := mostRepeated(input)
	fmt.Println(output)
}
