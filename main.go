package main

import (
	"fmt"
)

func main() {

	var degrees int
	fmt.Scan(&degrees) // считаем переменную 'a' с консоли

	hours := degrees / 30
	mins := 2 * (degrees % 30)
	message := fmt.Sprintf("It is %d hours %d minutes.", hours, mins)

	fmt.Println(message)
}
