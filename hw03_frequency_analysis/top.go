package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordsCount struct {
	word  string
	count int
}

var re = regexp.MustCompile(`[,.!?:;"'()\s]+-*\s*`)

func Top10(text string) []string {
	woCoMap := make(map[string]int)
	woCoSls := make([]wordsCount, 0, len(woCoMap))
	fin := 10
	ans := make([]string, 0, fin)
	if len(text) == 0 {
		return ans
	}
	unSignedText := re.Split(text, -1)
	for _, word := range unSignedText { // counts the number of occurrences of words
		woCoMap[strings.ToLower(word)]++
	}
	for word, count := range woCoMap {
		woCoSls = append(woCoSls, wordsCount{word, count})
	}
	sort.Slice(woCoSls, func(i, j int) bool { // word sorting
		if woCoSls[i].count > woCoSls[j].count {
			return true
		}
		if woCoSls[i].count < woCoSls[j].count {
			return false
		}
		return woCoSls[i].word < woCoSls[j].word
	})
	if len(woCoSls) < 10 {
		fin = len(woCoSls)
	}
	for i := 0; i < fin; i++ {
		ans = append(ans, woCoSls[i].word)
	}
	return ans
}
