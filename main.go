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

	if len(a) != 6 {
		panic("Ticket id must contains 6 digits")
	}

	if a[0]+a[1]+a[2] == a[3]+a[4]+a[5] {
		fmt.Println("YES")

		return
	}

	fmt.Println("NO")
}
