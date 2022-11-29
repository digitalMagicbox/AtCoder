package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d\n", &a)
	c := 0
	if a > 99 {
		c++
		a -= 100
	}
	if a > 9 {
		c++
		a -= 10
	}
	c += a

	fmt.Printf("%d\n", c)
}
