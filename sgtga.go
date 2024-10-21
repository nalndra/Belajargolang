package main

import "fmt"

func main() {
	var x, n, al, ting int
	var luas float64
	fmt.Scan(&x)

	for n = 1; n <= x; n += 1 {
		fmt.Scan(&al, &ting)
		luas = 0.5 * float64(al*ting)
		fmt.Print(luas)
	}
}
