package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanln(&input)
	var a []int

	for input > 0 {
		v := input % 10
		a = append(a, v)
		input /= 10
	}

	fmt.Println(a[len(a)-1])
}
