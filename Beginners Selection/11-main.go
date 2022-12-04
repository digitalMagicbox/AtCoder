package main

import (
	"fmt"
)

func move(s int, e int) int {
	m := e - s
	if m < 0 {
		return m * -1
	}
	return m
}

func main() {
	var n int

	fmt.Scan(&n)

	var t, x, y, mt, mx, my int
	for i := 0; i < n; i++ {
		fmt.Scan(&t, &x, &y)
		mt = t - mt
		m := move(mx, x) + move(my, y)

		if m == 0 {
			fmt.Print("No\n")
			return
		}
		if m > mt {
			fmt.Print("No\n")
			return
		}
		if m%2 != mt%2 {
			fmt.Print("No\n")
			return
		}
		mt = t
		mx = x
		my = y
	}
	fmt.Print("Yes\n")
}
