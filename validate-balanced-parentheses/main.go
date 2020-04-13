package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "([()])"
	s := strings.Split(a, "")
	res := true
	mapping := map[string]string{"{": "}", "(": ")", "[": "]"}

	if len(s)%2 != 0 {
		res = false
	}

	m := len(s)
	skip := make(map[int]bool)
	for i := 0; i < m; i++ {
		if !res {
			break
		}

		if skip[i] {
			continue
		}

		skipCounterpart := 0
		for j := i + 1; j < m; j++ {
			if s[i] == s[j] {
				skipCounterpart++
			}

			if mapping[s[i]] == s[j] {
				if skipCounterpart == 0 {
					if (j-i)%2 == 0 {
						res = false
						break
					}

					skip[j] = true
					break
				} else {
					skipCounterpart--
				}
			}

			// last string
			if mapping[s[i]] != s[j] && j == m-1 {
				res = false
				break
			}
		}
	}

	fmt.Println(res)
}
