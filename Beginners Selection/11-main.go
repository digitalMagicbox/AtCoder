package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func readLine() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := readLine()

	var mt, mx, my int
	for i := 0; i < n; i++ {
		t, x, y := readLine(), readLine(), readLine()
		mt = t - mt
		m := absInt(x-mx) + absInt(y-my)

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
