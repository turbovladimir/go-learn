package main

import "fmt"

func main() {
	//var a int
	m := make(map[int]int)
	c := 0

	for i := 0; i < 10; i++ {
		fmt.Scan(&c)

		v, ok := m[c]

		if !ok {
			v = work(c)
			m[c] = v
		}

		fmt.Print(v)
		fmt.Print(" ")
	}
}
func work(i int) int {
	return i
}
