package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "aabceabdfvc"
	s := strings.Split(input, "")
	n := len(s)
	maxLen := 1
	curLen := 1
	var prevIdx int
	visited := make(map[string]int)

	visited[s[0]] = 0
	for i := 1; i < n; i++ {
		prevIdx = visited[s[i]]

		if prevIdx == -1 || i-curLen > prevIdx {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			// current idx - last seen
			curLen = i - prevIdx
		}
		// last seen
		visited[s[i]] = i
	}

	if curLen > maxLen {
		maxLen = curLen
	}

	fmt.Println(maxLen)
}
