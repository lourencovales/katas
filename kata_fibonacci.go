/**
The Fibonacci numbers are the numbers in the following integer sequence (Fn):

    0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...

such as

    F(n) = F(n-1) + F(n-2) with F(0) = 0 and F(1) = 1.

Given a number, say prod (for product), we search two Fibonacci numbers F(n) and F(n+1) verifying

    F(n) * F(n+1) = prod.

Your function productFib takes an integer (prod) and returns an array:

[F(n), F(n+1), true] or {F(n), F(n+1), 1} or (F(n), F(n+1), True)

depending on the language if F(n) * F(n+1) = prod.

If you don't find two consecutive F(n) verifying F(n) * F(n+1) = prodyou will return

[F(n), F(n+1), false] or {F(n), F(n+1), 0} or (F(n), F(n+1), False)

F(n) being the smallest one such as F(n) * F(n+1) > prod.
**/

package main

import (
	"fmt"
)

func main() {
	a := ProductFib(74049690)
	fmt.Printf("%d\n", a)
}

// We iterate until we find that we've either went past the product value, or that it's equal to it

func ProductFib(prod uint64) [3]uint64 {

	var k [3]uint64
	for i := 0; i < int(prod); i++ {
		j := uint64(i)
		if uint64(Product(j)) > prod {
			k = [3]uint64{Fib(j), Fib(j + 1), 0}
			break
		} else if uint64(Product(j)) == prod {
			k = [3]uint64{Fib(j), Fib(j + 1), 1}
			break
		}
	}
	return k
}

// Simple product of the Fib values

func Product(n uint64) uint64 {
	return Fib(n) * Fib(n+1)
}

// We calculate Fib by storing the past values and iterating over; more efficient

func Fib(n uint64) uint64 {
	a := make([]uint64, n+2)
	a[0], a[1] = 0, 1
	for i := 2; i <= int(n); i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
}
