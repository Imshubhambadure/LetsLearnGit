package main

import "fmt"

func main() {
	var weight int
	fmt.Scan(&weight)

	if weight%2 == 0 && weight <= 100 && weight > 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
