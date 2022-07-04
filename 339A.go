package main

import (
	"fmt"
)

func main() {
	var a [100]string
	fmt.Scan(&a)
	//lth := len(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] != "+" {
				if a[j] > a[j+2] {
					var temp string
					temp = a[j]
					a[j] = a[j+2]
					a[j+2] = temp
				}
			}
		}
	}
	fmt.Printf("%s", a)
}
