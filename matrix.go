package main

import "fmt"

func main() {
	var mat [5][5]int
	var sum, count, m, n, k int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&mat[i][j])
			sum = sum + mat[i][j]
			if sum == 1 {
				k++
				if k == 1 {
					m = i
					n = j
				}
			}
		}
	}
	for n != 2 {
		count++
		if n >= 3 {
			n--
		} else {
			n++
		}
	}
	for m != 2 {
		count++
		if m >= 3 {
			m--
		} else {
			m++
		}
	}
	fmt.Printf("%d", count)
}
