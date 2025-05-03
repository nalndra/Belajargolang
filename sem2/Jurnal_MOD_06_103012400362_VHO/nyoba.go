package main
import "fmt"

func main(){
	var n int
	
	fmt.Scan(&n)
	fmt.Println(fibofiboan(n))
	
}

func fibofiboan(n int)int {
	if n == 1 || n == 2{
		return 1
	}
	if n == 3{
		return 4
	}else{
		return fibofiboan(n-1)+fibofiboan(n-2)+fibofiboan(n-3)
	
	}
}