package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	s := strings.Split(a, "")
	nums := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if s[i] != "+" {
			nums = append(nums, s[i])
		}
	}
	sort.Strings(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%s", nums[i])
		if i < len(nums)-1 {
			fmt.Print("+")
		}
	}
}
