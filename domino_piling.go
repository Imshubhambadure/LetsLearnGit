package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	s := m * n
	fmt.Printf("%d\n", s/2)
}
