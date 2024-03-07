package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 2}, {3, 5}},
		{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
	}

	for _, test := range tests {
		fmt.Println("case:", test)
		fmt.Println("result:", merge(test))
		fmt.Println()
	}
}

func merge(intervals [][]int) [][]int {
	// sort interval base on intervals[i][0]
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	results := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if i > 0 && isConflict(results[len(results)-1], intervals[i]) {
			results[len(results)-1] = mergeInterval(results[len(results)-1], intervals[i])
		} else {
			results = append(results, intervals[i])
		}
	}

	return results
}

func mergeInterval(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func isConflict(a, b []int) bool {
	return a[1] >= b[0]
}
