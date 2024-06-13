package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// Errors

var ErrEmptyInput = errors.New("no words in empty string")

// Regex explanation:
// (?im) -> Case insensitive match of characters, multiline
// [a-zа-я0-9]+ -> Allows words with at least 1 character in it
// (?:[.,!?@'"_\-]+[a-zа-я0-9]+)* -> Allows at 1 or multiple special character between regular characters
// (?:[\-]{2,}) -> Allows words with at least 2 allowed special characters

var re = regexp.MustCompile(`(?im)[a-zа-я0-9]+(?:[.,:;!?@'"_\-]+[a-zа-я0-9]+)*|(?:[\-]{2,})`)

// Custom types

type WordFrequency struct {
	Word      string
	Frequency int
}

type WordFrequencyList []WordFrequency

func (wfl WordFrequencyList) Len() int {
	return len(wfl)
}

func (wfl WordFrequencyList) Less(i, j int) bool {
	if wfl[i].Frequency == wfl[j].Frequency {
		// If frequency is equal, then sort words lexicographically (ascending)
		return strings.Compare(wfl[i].Word, wfl[j].Word) > 0
	}
	return wfl[i].Frequency < wfl[j].Frequency
}

func (wfl WordFrequencyList) Swap(i, j int) {
	wfl[i], wfl[j] = wfl[j], wfl[i]
}

func (wfl WordFrequencyList) TopFrequencies(n int) WordFrequencyList {
	sort.Sort(sort.Reverse(wfl))
	if n >= len(wfl) {
		return wfl
	}
	return wfl[:n]
}

func (wfl WordFrequencyList) TopWords(n int) []string {
	rv := make([]string, 0)
	tfl := wfl.TopFrequencies(n)
	for _, v := range tfl {
		rv = append(rv, v.Word)
	}
	return rv
}

// Functions

func FrequencyAnalisys(text string) WordFrequencyList {
	words := re.FindAllString(text, -1)
	m := make(map[string]int)
	for _, w := range words {
		m[strings.ToLower(w)]++
	}
	wfl := make(WordFrequencyList, 0)
	for word, frequency := range m {
		wfl = append(wfl, WordFrequency{word, frequency})
	}
	return wfl
}

func Top10(text string) []string {
	return FrequencyAnalisys(text).TopWords(10)
}
