package main

import "fmt"

func main() {
	var w int
	var v, pos float64
	fmt.Scan(&w, &v)

	for i := 0; i <= w; i++ {
		pos = (v * float64(i)) + (0.5 * 9.8 * float64(i) * float64(i))
		fmt.Printf("%.5f\n", pos)
	}
}
