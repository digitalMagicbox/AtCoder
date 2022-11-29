package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)
	min := -1
	for i := 0; i < n; i++ {
		for j := 1; ; j++ {
			if min < j && min >= 0 {
				break
			}
			if x[i]%2 == 1 {
				if j <= min || min == -1 {
					min = j - 1
				}
				break
			}
			x[i] = x[i] / 2
		}
	}

	fmt.Printf("%d\n", min)
}
