package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}
	sort.Ints(c)
	var a, b int
	for i := n - 1; 0 <= i; i-- {
		a += c[i]
		i--
		if i == -1 {
			break
		}
		b += c[i]
	}

	fmt.Println(a - b)
}
