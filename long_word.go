package main

import "fmt"

func main() {

	var n int
	fmt.Scanln(&n)
	var names = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&names[i])
	}
	for i := 0; i < n; i++ {
		if len(names[i]) > 10 && len(names[i]) < 100 {
			fmt.Printf("INVALID")
		} else {
			fmt.Printf("%s\n", names[i])
		}

	}
}
