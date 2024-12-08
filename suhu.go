package main
import "fmt"
func main(){
	var cel, fah, suhu int
	var temp string

	fmt.Scan(&suhu, &temp)
	if temp == "Celcius"{
		fah = ((9 * suhu) / 5 ) + 32
		fmt.Printf("Suhu dalam fahrenheit adalah %d F",fah )
	}else if temp == "Fahrenheit" {
		cel = (suhu - 32) * 5 / 9
		fmt.Printf ("Suhu dalam celcius adalah %d C", cel)
	}
}
