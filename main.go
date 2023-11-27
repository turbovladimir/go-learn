package main

import "fmt"

func main() {
	var n, s, a int

	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		a = arr[i]

		if a/10 < 10 && a/10 > 0 && a%8 == 0 {
			s += a
		}
	}

	println(s)
}
