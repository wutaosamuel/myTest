package main

import (
	"sort"
)

// SortUnityIndexes sort index and deduplicate
func SortUnityIndexes(indexes []int, maxIndex int) []int {
	// sort indexes slice
	sort.Ints(indexes)
	// remove duplicate && lower index && higher index
	deduplicate := -1
	index := make([]int, 0)
	for _, v := range indexes {
		if v < 0 {
			continue
		}
		if deduplicate == v {
			continue
		}
		if v > maxIndex {
			break
		}
		deduplicate = v
		index = append(index, v)
	}

	return index
}
