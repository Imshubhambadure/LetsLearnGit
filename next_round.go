package main

import (
	"fmt"
)

func main() {
	var n, k, f int
	fmt.Scanln(&n, &k)
	var round = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&round[i])
	}
	score := round[k-1]
	for j := 0; j < n; j++ {
		if round[j] >= score && round[j] != 0 {
			f++
		}
	}
	fmt.Printf("%d", f)
}
