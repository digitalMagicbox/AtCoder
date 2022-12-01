// ABC085B - Kagami Mochi
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	m := make([]int, n)
	fmt.Scan(&a)
	m[0] = a
	c := 1
	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		j := 0
		for ; j < c; j++ {
			if m[j] == a {
				break
			}
		}
		m[j] = a
		if c <= j {
			c = j + 1
		}
	}

	fmt.Println(c)
}
