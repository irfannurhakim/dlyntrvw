package main

import "fmt"

func main() {
	s := []int{3, 3, 2, 1, 3, 2, 1}
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			i = -1
		}
	}
	fmt.Println(s)
}
