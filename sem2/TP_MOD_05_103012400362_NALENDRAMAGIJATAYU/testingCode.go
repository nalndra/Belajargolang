package main
import "fmt"

func fibonacci(n int) int {
    if n == 1 || n == 2 {
        return 1
    }else{
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var hasil int
    var bil int
	
    fmt.Scan(&bil)

	hasil = fibonacci(bil)
	
    fmt.Println(hasil)
}
