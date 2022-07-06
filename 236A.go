package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	//str := make([]string, 0)
	b := strings.Split(a, "")
	// fmt.Println(len(b))
	// var count = 1
	// for i := 0; i < len(a)-1; i++ {
	// 	if a[i] != a[i+1] {
	// 		count++
	// 	}

	// }
	var ds = make(map[string]bool)
	var unique_ls []string
	for _, val := range b {
		if _, ds_val := ds[val]; !ds_val {
			ds[val] = true
			unique_ls = append(unique_ls, val)
		}
	}
	if len(unique_ls)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
