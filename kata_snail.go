package main

import "fmt"

func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	final := Snail(matrix)
	fmt.Println(final)

}

func Snail(snaipMap [][]int) []int {
	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return []int{}
	}
	lenght := len(snaipMap)
	t, b, l, r := 0, lenght-1, 0, lenght-1
	final := make([]int, 0, lenght*lenght)

	for t <= b && l <= r {
		for i := l; i <= r; i++ {
			final = append(final, snaipMap[t][i])
		}
		t++
		for i := t; i <= b; i++ {
			final = append(final, snaipMap[i][r])
		}
		r--
		if t <= b {
			for i := r; i >= l; i-- {
				final = append(final, snaipMap[b][i])
			}
			b--
		}
		if l <= r {
			for i := b; i >= t; i-- {
				final = append(final, snaipMap[i][l])
			}
			l++
		}
	}
	return final
}
