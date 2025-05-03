package main
import "fmt"

func faktorial(n int) int {
    if n == 1 { 
        return 1
    }else{
		return n * faktorial(n - 1) 
	}
}

func main() {
    var x, hasil int
    fmt.Scan(&x)

    hasil = faktorial(x)
	
    fmt.Println(hasil)
}
