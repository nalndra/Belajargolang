package main

import "fmt"
func main(){
	var ktp, surat bool
	fmt.Scan(&ktp, &surat)

	if ktp{
		fmt.Print("Diizinkan masuk")
	}else if surat {
		fmt.Print("Diizinkan masuk") 
	}else {
		fmt.Print("Tidak diizinkan masuk")
	}
}