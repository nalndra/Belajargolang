package main

import "fmt"

func segitiga(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 8 {
		return false
	}
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func main() {
	var a, b, c float64
	fmt.Print("Masukkan panjang sisi a b c: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Penjelasan: Masa sisi segitiga negatif")
		return
	}

	if a > b+c {
		fmt.Printf("Penjelasan: (%.2f) terlalu panjang dibandingkan yang lain\n", a)
		return
	} else if b > a+c {
		fmt.Printf("Penjelasan: (%.2f) terlalu panjang dibandingkan yang lain\n", b)
		return
	} else if c > a+b {
		fmt.Printf("Penjelasan: (%.2f) terlalu panjang dibandingkan yang lain\n", c)
		return
	}

	if segitiga(a, b, c) {
		fmt.Printf("%.2f, %.2f, dan %.2f Segitiga? true\n", a, b, c)
		fmt.Println("Penjelasan: Ya iyalah")
	} else {
		fmt.Printf("%.2f, %.2f, dan %.2f Segitiga? false\n", a, b, c)
	}
}
