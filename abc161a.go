package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	sum := 0

	sum += x / 500 * 1000
	sum += (x % 500) / 5 * 5

	fmt.Println(sum)
}