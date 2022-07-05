package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	//str := make([]string, 0)
	var count = 1
	for i := 0; i < len(a)-1; i++ {
		if a[i] != a[i+1] {
			count++
		}

	}
	if count%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
