package main

import (
	"fmt"
	"strings"
)

func main() {
	t := []string{"dream", "dreamer", "erase", "eraser"}

	var s string
	fmt.Scan(&s)
	for {
		i := 0
		for ; i < len(t); i++ {
			if strings.HasSuffix(s, t[i]) {
				s = s[:len(s)-len(t[i])]
				if len(s) == 0 {
					fmt.Print("YES\n")
					return
				}
				break
			}
		}
		if i == len(t) {
			break
		}
	}
	fmt.Print("NO\n")
}
