//Given an array of integers, find the one that appears an odd number of times.
//
//There will always be only one integer that appears an odd number of times.

package main

import "fmt"

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
}

func FindOdd(seq []int) int {
	array := make(map[int]int)
	for _, v := range seq {
		array[v] += 1
	}
	for i, x := range array {
		if x%2 != 0 {
			return i
		}
	}
	return 0 // your code here
}
