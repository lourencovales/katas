/**
A Hamming number is a positive integer of the form 2i3j5k, for some non-negative integers i, j, and k.

Write a function that computes the nth smallest Hamming number.

Specifically:

    The first smallest Hamming number is 1 = 203050
    The second smallest Hamming number is 2 = 213050
    The third smallest Hamming number is 3 = 203150
    The fourth smallest Hamming number is 4 = 223050
    The fifth smallest Hamming number is 5 = 203051

**/

package main

import "fmt"

func main() {
	fmt.Println(Hammer(282))
}

func Hammer(n int) int {
	h := make([]int, n)
	x2, x3, x5 := 2, 3, 5
	i, j, k := 0, 0, 0

	if n == 1 {
		return 1
	}

	for u := 1; u < n; u++ {
		nh := min(x2, x3, x5)
		h[u] = nh
		if x2 == nh {
			i += 1
			x2 = 2 * h[i]
		}
		if x3 == nh {
			j += 1
			x3 = 3 * h[j]
		}
		if x5 == nh {
			k += 1
			x5 = 5 * h[k]
		}
	}
	return h[n-1]
}
