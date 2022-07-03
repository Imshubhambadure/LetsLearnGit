package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var b string
	fmt.Scan(&a, &b)
	var s1, s2 string
	for i := 0; i < len(a); i++ {
		s1 = strings.ToLower(a)
		s2 = strings.ToLower(b)
	}
	var count int
	if s1 < s2 {
		count = -1
	}
	if s1 > s2 {
		count = 1
	}
	if s1 == s2 {
		count = 0
	}
	fmt.Printf("%d", count)
}
