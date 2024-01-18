package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		// "a" harfi sayısına göre azalan sıra, eğer aynı sayıda ise uzunluğa göre sırala
		countA := strings.Count(words[i], "a") > strings.Count(words[j], "a")
		length := len(words[i]) > len(words[j])

		return countA || (countA && length)
	})
	return words
}

func main() {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	output := sortByA(input)
	fmt.Println(output)
}
