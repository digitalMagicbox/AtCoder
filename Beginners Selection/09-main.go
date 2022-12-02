package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)

	for x := N; x >= 0; x-- {
		for y := 0; y < (N - x + 1); y++ {
			z := (N - x - y)
			if Y == x*10000+y*5000+z*1000 {
				fmt.Printf("%d %d %d\n", x, y, z)
				return
			}
		}
	}

	fmt.Printf("-1 -1 -1\n")
}
