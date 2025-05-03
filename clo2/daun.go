package main

import "fmt"

func main() {
	var k, total, jumlah, min, max int
	var r float64

	min = 101
	max = -1

	for {
		fmt.Scan(&k)

		if k == 100 {
			break
		} else if k >= 1 && k <= 100 {
			total++
			jumlah += k
			if k < min {
				min = k
			}
			if k > max {
				max = k
			}
		}
	}
	r = float64(jumlah / total)
	fmt.Printf("%d %.2f %d %d", total, r, min, max)
}