package main

import (
	"fmt"
)

func main() {
	str := "tracecars"
	maxLen := 0

	for i := 0; i <= len(str); i++ {
		for j := i; j <= len(str); j++ {
			s := str[i:j]
			if len(s) > maxLen && isPolindrom(s) {
				fmt.Println(s)
				maxLen = len(s)
			}
		}

	}

	fmt.Println(maxLen)
}

func isPolindrom(str string) bool {
	if str == reverse(str) {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
