package main

import (
	"fmt"
)

func main() {
	var input int32
	fmt.Scanln(&input)

	switch {
	case input > 0:
		fmt.Println("Число положительное")
	case input < 0:
		fmt.Println("Число отрицательное ")
	case input == 0:
		fmt.Println("Ноль")
	}
}
