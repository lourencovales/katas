package main

import (
	"fmt"
)

func main() {
	a := Race(720, 850, 70)
	fmt.Print(a)
}

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	t := float64(g) / float64(v2-v1)

	hours := int(t)
	minutes := int((t - float64(hours)) * 60)
	seconds := int((t - float64(hours) - float64(minutes)/60) * 3600)

	return [3]int{hours, minutes, seconds}

}
