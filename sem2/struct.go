package main
import "fmt"

type segiEmpat struct {
	sisi, kel, luas float64
}

func main (){
	var n int
	var persegi segiEmpat

	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		fmt.Scan(&persegi.sisi)
		persegi.kel = 4 * persegi.sisi
		persegi.luas = persegi.sisi * persegi.sisi

		fmt.Println(persegi.kel, persegi.luas)
	}
}