package main
import "fmt"

func main(){
	var n, i int
	var nama string
	var status string
	var angkatan int
	
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		inpData(&nama, &status, &angkatan)
		if i == 0{
			fmt.Println("                    DATA")
			fmt.Println("__________________________________________________")
		}
		outputData(nama, status, angkatan)
	}
}

func inpData(nama *string, status *string, angkatan *int){
	fmt.Scan(nama, status, angkatan)
}

func outputData(nama, status string, angkatan int){
	fmt.Printf("| %-15s | %-20s | %-5d |\n", nama, status, angkatan)
}
