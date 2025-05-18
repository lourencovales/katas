package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("vim-go")
}

func Mix(s1, s2 string) string {
	for _, r := range s1 {
		if unicode.IsLower(r) {
			news1 += r
		}
	}
	for _, r := range s2 {
		if unicode.IsLower(r) {
			news2 += r
		}
	}
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	mapping1 := make(map[rune]int)
	mapping2 := make(map[rune]int)

	for _, r := range lowercase {
		count := strings.Count(news1, r)
		mapping1[r] = count
	}
	for _, r := range lowercase {
		count := strings.Count(news2, r)
		mapping2[r] = count
	}

	var mapping3 make(map[rune]int)

	for k, v := range mapping1 {
		for k2, v2 := range mapping2 {
			if k == k2 && v2 > v {
				mapping3[k] = v2
}
		}
	}
}
