package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ValidISBN10("1112223339"))
}

func ValidISBN10(isbn string) bool {
	var isbnSlice []int
	var sum int
	if len(isbn) != 10 {
		return false
	}
	for _, c := range isbn {
		intC, err := strconv.Atoi(string(c))
		if err != nil {
			if (string(c) == "X" || string(c) == "x") && len(isbnSlice) == 9 {
				intC = 10
			} else {
				return false
			}
		}
		isbnSlice = append(isbnSlice, intC)
	}
	for i, n := range isbnSlice {
		sum += (i + 1) * n
	}

	return sum%11 == 0
}
