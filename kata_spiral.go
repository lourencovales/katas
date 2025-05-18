package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func Spiralize(size int) [][]int {
	mapping := make([][]int, 0, size*size)
	final := make([][]int, 0, size*size)
	x, y := 0, 0

	for x = 0; x < size-1; x++ {
		for y = 0; y < size-1; y++ {
			mapping[x][y] = 0
		}
	}

}
