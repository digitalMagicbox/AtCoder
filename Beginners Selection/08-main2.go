package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		m[a] = struct{}{}
	}
	fmt.Println(len(m))
}
