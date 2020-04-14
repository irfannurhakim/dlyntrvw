package main

import (
	"fmt"
)

func main() {
	s := []int{3, 3, 5, 7, 8, 9, 9, 9}
	target := 9

	n := len(s)
	minIdx := 0
	maxIdx := 0
	found := false // found duplicate occurences

	for i := 0; i < n; i++ {
		if s[i] == target {
			if i == 0 || s[i-1] != target {
				minIdx = i
			} else {
				maxIdx = i
				found = true
			}
		}
	}

	if !found {
		minIdx = -1
		maxIdx = -1
	}

	fmt.Printf("start: %d end: %d", minIdx, maxIdx)
}
