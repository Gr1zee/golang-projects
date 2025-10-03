package main

import (
	"fmt"
	"slices"
)

func main() {
	m := map[string]int{
		"Alice":   10,
		"Bob":     1000,
		"Charlie": 1,
	}
	fmt.Println(getTopWords(m, 2))
}

func AnalyzeText(text string) {

}

func getTopWords(wordMap map[string]int, n int) []string {
	var counts []int
	topWords := make([]string, n)

	for _, v := range wordMap {
		counts = append(counts, v)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	for i := 0; i < n; i++ {
		for k, v := range wordMap {
			if counts[i] == v && !slices.Contains(topWords, k) {
				topWords[i] = k
			}
		}
	}
	return topWords
}
