package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)

	y = 0

	for x > 0 {
		x = x / 10
		y += 1
	}
	fmt.Printf("%d digit", y)
}
