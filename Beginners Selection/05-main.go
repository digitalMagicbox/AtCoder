package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	y := 0
	for i := 0; i <= a; i++ {
		if x < 500*i {
			break
		}
		for j := 0; j <= b; j++ {
			if x < 500*i+100*j {
				break
			}
			for k := 0; k <= c; k++ {
				if x == 500*i+100*j+50*k {
					y++
					break
				} else if x < 500*i+100*j+50*k {
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", y)
}
