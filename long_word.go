package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var names = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&names[i])
	}
	for i := 0; i < n; i++ {
		if len(names[i]) > 10 && len(names[i]) < 100 {
			long_word := names[i]
			// fmt.Print(string((len(long_word)) - 2))
			fmt.Println(long_word[0:1] + strconv.Itoa(len(long_word)-2) + long_word[len(long_word)-1:])
		} else {
			fmt.Printf("%s\n", names[i])
		}

	}
}
