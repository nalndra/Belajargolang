package main

import "fmt"

func kesbek(kode int) bool {
	bilPertama := kode / 1000
	bilEmpat := kode % 10
		return bilPertama == bilEmpat
}
func main() {
	var kode int
	fmt.Print("Periksa Cashback: ")
	fmt.Scan(&kode)
	fmt.Println(kesbek(kode))
}
