package main

import (
	"fmt"
	"slices"
)

func main() {
	var input int
	fmt.Scanln(&input)
	var a []int

	for input > 0 {
		v := input % 10

		if slices.Contains(a, v) {
			fmt.Println("NO")
			return
		}

		a = append(a, v)
		input /= 10
	}

	fmt.Println("YES")
}
