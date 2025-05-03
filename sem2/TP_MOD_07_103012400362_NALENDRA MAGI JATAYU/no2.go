package main
import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
		return
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(A[i])
	}
	fmt.Println()
}

func main() {
	var A tabInt
	var n int

	bacaData(&A, &n)
	cetakData(A, n)
}
