/*

The prime numbers are not regularly spaced. For example from 2 to 3 the gap is 1. From 3 to 5 the gap is 2. From 7 to 11 it is 4. Between 2 and 50 we have the following pairs of 2-gaps primes: 3-5, 5-7, 11-13, 17-19, 29-31, 41-43

A prime gap of length n is a run of n-1 consecutive composite numbers between two successive primes (see: http://mathworld.wolfram.com/PrimeGaps.html).

We will write a function gap with parameters:

    g (integer >= 2) which indicates the gap we are looking for

    m (integer > 2) which gives the start of the search (m inclusive)

    n (integer >= m) which gives the end of the search (n inclusive)

In the example above gap(2, 3, 50) will return [3, 5] or (3, 5) or {3, 5} which is the first pair between 3 and 50 with a 2-gap.

So this function should return the first pair of two prime numbers spaced with a gap of g between the limits m, n if these numbers exist otherwise `nil or null or None or Nothing (or ... depending on the language). 

*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// We take args from stdin, convert the arg from string to int, and use it as input for the checking function

func main() {
	f, _ := strconv.Atoi(os.Args[1])
	s, _ := strconv.Atoi(os.Args[2])
	t, _ := strconv.Atoi(os.Args[3])
	result := gap(f, s, t)
	if result != nil {
		fmt.Printf("%v", result)
	} else {
		fmt.Println("Error")
	}
}

// This is a function to check if the number is prime

func isPrime(a int) bool {

	if a <= 1 {
		return false
	} else if a == 2 {
		return true
	} else if a%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(a)))
	for i := 3; i <= sqrt; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
}

// The problems calls for "consecutive primes", so we need to check if there are any primes in between the ones we're checking

func containsPrime(a, b int) bool {
	for i := a + 1; i < b; i++ {
		if isPrime(i) {
			return false
		}
	}
	return true
}

// This is the main gap checking function

func gap(g, m, n int) []int {
	var interval []int
	for i := m; i <= n; i++ {
		o := i + g
		if isPrime(i) && isPrime(o) && o <= n && containsPrime(m, n) {
			interval := append(interval, i, o)
			return interval
		}
	}
	return nil
}
