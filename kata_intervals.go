package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

func SumOfIntervals(intervals [][2]int) int {
	mergedintervals := [][2]int{}

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	current := intervals[0]

	for _, interval := range intervals[1:] {
		if interval[0] <= current[1] {
			if interval[1] > current[1] {
				current[1] = interval[1]
			}

		} else {
			mergedintervals = append(mergedintervals, current)
			current = interval
		}
	}
	mergedintervals = append(mergedintervals, current)

	lenght := 0
	for _, interval := range mergedintervals {
		lenght += interval[1] - interval[0]
	}
	return lenght
}
