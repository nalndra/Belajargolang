package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var hasil bool = true
	
	for i := 2; i * i <= x; i++{
		if x % i == 0{
			hasil = false
		} 
	}
	fmt.Println(hasil)
}