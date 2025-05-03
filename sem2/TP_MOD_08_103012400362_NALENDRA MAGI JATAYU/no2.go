package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(sumNumSeq(n))
}

func numSeq(n int)int{
	if n == 1 || n == 0{
		return 1
	} else{
		return n * numSeq(n-1)
	}
}

func sumNumSeq(n int)int{
	if n == 1{
		return numSeq(1)
	}else {
		return numSeq(n) + sumNumSeq(n-1)
	}
}
