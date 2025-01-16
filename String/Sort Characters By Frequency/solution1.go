package main

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   rune
	Value int
}

type ByPair []Pair

func (p ByPair) Len() int           { return len(p) }
func (p ByPair) Less(i, j int) bool { return p[i].Value > p[j].Value } // Descending order
func (p ByPair) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func frequencySort(s string) string {
	freqMap := make(map[rune]int)
	for _, c := range s {
		freqMap[c]++
	}
	items := make([]Pair, 0, len(freqMap)) // Initialize with map size
	for k, v := range freqMap {
		items = append(items, Pair{k, v})
	}
	sort.Sort(ByPair(items)) // Sort directly in descending order

	var builder strings.Builder
	for _, item := range items {
		builder.WriteString(strings.Repeat(string(item.Key), item.Value))
	}
	return builder.String()
}

func main() {
	print(frequencySort("tree"))
}
