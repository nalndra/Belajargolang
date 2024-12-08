package main

import "fmt"

func main () {
	var m, n int
	fmt.Scan (&m, &n)

	baris := make([]int, n)
	for i := 0; i <n; i++ {
	fmt.Scan(&baris[1])
	}

	benar := true
	for i := range baris {
			benar = benar && (baris[1] == m*(i+1))
	}

	fmt.Println(benar)
}