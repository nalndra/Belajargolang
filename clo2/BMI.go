package main
import "fmt"
func main(){
	var tb, bb, bmi float64
	fmt.Scan(&tb, &bb)

	bmi = bb / (tb/100*tb/100)

	if bmi >= 18.5 && bmi <= 22.9{
		fmt.Print("Normal")
	}else if bmi > 22.9{
		fmt.Print("Berlebihan")
	}else if bmi < 18.5{
		fmt.Print("Kekurangan")
	}
}

//gitu yhhhh