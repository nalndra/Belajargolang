package main

import "fmt"

func thnKbs(thn int) bool {
	if thn%4 == 0 {
		return true
	}
	if thn%100 == 0 {
		return false
	}
	return false
}
func main(){
	var thn int
	fmt.Print("Periksa tahun kabisat: ")
	fmt.Scan(&thn)
	fmt.Println(thnKbs(thn))
}