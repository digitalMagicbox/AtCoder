package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Scan(&s)

	s = strings.ReplaceAll(s, "eraser", "")
	s = strings.ReplaceAll(s, "erase", "")
	s = strings.ReplaceAll(s, "dreamer", "")
	s = strings.ReplaceAll(s, "dream", "")

	if len(s) == 0 {
		fmt.Print("YES\n")
	} else {
		fmt.Print("NO\n")
	}
}
