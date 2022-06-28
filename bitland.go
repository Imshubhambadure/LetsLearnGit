package main

import (
	"fmt"
)

func main() {
	var n, f int
	fmt.Scan(&n)
	var x string
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x[0] == '-' || x[1] == '-' {
			f--
		}
		if x[0] == '+' || x[1] == '+' {
			f++
		}
	}
	fmt.Printf("%d\n", f)
}
