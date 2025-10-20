package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
}

func AnalyzeText(text string) {
	newText := strings.ReplaceAll(text, ".", " ")
	newText = strings.ReplaceAll(newText, ",", " ")
	newText = strings.ReplaceAll(newText, "!", " ")
	newText = strings.ReplaceAll(newText, "?", " ")
	words := strings.Split(newText, " ")
	counts := make(map[string]int)
	totalCountWords := 0
	uniqueCountWords := 0

	for _, word := range words {
		if word != "" {
			counts[strings.ToLower(word)]++
			totalCountWords++
		}
	}
	sortredKeys := make([]string, 0, len(counts))
	for k := range counts {
		sortredKeys = append(sortredKeys, k)
		uniqueCountWords++
	}

	res := getTopWords(counts, 5)
	fmt.Printf("Количество слов: %d\n", totalCountWords)
	fmt.Printf("Количество уникальных слов: %d\n", uniqueCountWords)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", res[0], counts[res[0]])
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	fmt.Printf("\"%s\": %d раз\n", res[0], counts[res[0]])
	fmt.Printf("\"%s\": %d раз\n", res[1], counts[res[1]])
	fmt.Printf("\"%s\": %d раз\n", res[2], counts[res[2]])
	fmt.Printf("\"%s\": %d раз\n", res[3], counts[res[3]])
	fmt.Printf("\"%s\": %d раз\n", res[4], counts[res[4]])
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
