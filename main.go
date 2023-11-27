package main

import "fmt"

func main() {
	var a, b, s int
	fmt.Scan(&a)
	fmt.Scan(&b)

	for a <= b {
		s += a
		a++
	}

	println(s)
}
