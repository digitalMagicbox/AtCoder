package main

import "fmt"

func main() {
	var n, a, b, total int
	fmt.Scan(&n, &a, &b)
	for i := 1; i <= n; i++ {
		sum := 0
		v := i
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		if a <= sum && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}
