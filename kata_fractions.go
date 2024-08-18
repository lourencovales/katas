package main

import "fmt"

func main() {
	fmt.Println(ProperFractions(25))
}

func ProperFractions(n int) int {
	if n == 1 {
		return 0
	}

	total := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n = n / i
			}
			total -= total / i
		}
	}
	if n > 1 {
		total -= total / n
	}
	return total
}
