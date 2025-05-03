package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	printNumRecursive(n)
}

func printNumRecursive(n int){
	if n > 0{
		printNumRecursive(n-1)
		fmt.Print(n, " ")
	}
}