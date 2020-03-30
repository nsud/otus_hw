package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
)

const sizeBox = 10

type countsWord struct {
	word  string
	count int
}

func Top10(text string) []string {
	if text == "" {
		return nil
	}
	var countArr []countsWord

	words := regexp.MustCompile(`[\s\t\r\n]+`).Split(text, -1)

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	defaultCount := 1
	oldWord := ""
	for _, word := range words {
		switch {
		case oldWord == word:
			{
				defaultCount++
			}
		case (oldWord != word) && (oldWord != ""):
			{
				countArr = append(countArr, countsWord{word: oldWord, count: defaultCount})
				defaultCount = 1
			}
		}
		oldWord = word
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i].count > countArr[j].count
	})

	top := []string{}

	for _, val := range countArr {
		top = append(top, val.word)
	}
	top10 := top[:sizeBox]

	return top10
}
