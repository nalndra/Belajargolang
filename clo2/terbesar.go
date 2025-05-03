package main

import "fmt"

func main() {
	var x, y, z, i int
	fmt. Scan(&x)

	y=0
	z=0
	for i = 1; i <= x; i ++ {
		fmt.Scan(&y)

		if y > z {
			z= y
		}
	}
	fmt.Print(z)
}