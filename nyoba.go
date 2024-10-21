package main

import "fmt"

func main(){
	var a, b, c, d int

	fmt.Print ("masukkan 4 digit kode: ")
	fmt.Scanf ("%d%d%d%d", &a, &b, &c, &d)

	if a == d {
		fmt.Print("true")
	}else {
		fmt.Print("false")
	}
}