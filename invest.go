package main

import "fmt"

func main() {
	var u float64
	fmt.Scan(&u)

	for t := 1; +t <= 10; t++ {
		u = u + (u * 0.05)
		fmt.Printf("Tahun ke-%d : Rp%2f\n", +t, u)
	}
}
