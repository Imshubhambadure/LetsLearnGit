package main

import (
	"fmt"
)

func main() {
	var n, x, y, z, f int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x, &y, &z)
		if (x + y + z) >= 2 {
			f++

		}

	}
	fmt.Printf("%d", f)
}
