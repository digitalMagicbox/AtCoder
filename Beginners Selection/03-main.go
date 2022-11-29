package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d\n", &a)
	b := (a % 100) / 10
	c := a % 10
	if a > 99 {
		a = 1
	} else {
		a = 0
	}
	fmt.Printf("%d\n", a+b+c)
}
