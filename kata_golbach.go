package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CheckGolbach(2000))
}

func CheckGolbach(n int) [2]int {
	var result [2]int
	var primes []int
	if n <= 2 || n%2 != 0 {
		return [2]int{}
	}
	for i := 2; i <= n; i++ {
		if !isPrime(i) {
			continue
		}
		primes = append(primes, i)
	}
	for j := len(primes) - 1; j >= 0; j-- {
		for i := 0; i < len(primes); i++ {
			if primes[i]+primes[j] == n {
				return [2]int{primes[i], primes[j]}
			}
		}
	}
	return result
}

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	if x == 2 || x == 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(x)))
	for i := 5; i <= sqrt; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
}
