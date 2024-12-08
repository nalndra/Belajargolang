package main
import "fmt"

func main() {
	var a int
	var spell string
	fmt.Scan(&a)

	switch{
	case a > 0 && a < 10:
		spell = "Satuan"
	case a >= 10 && a < 100:
		spell = "Puluhan"
	case a >= 100 && a < 1000:
		spell = "Ratusan"
	case a >= 1000 && a < 10000:
		spell = "Ribuan"
	case a >= 10000 && a < 100000:
		spell = "Puluhan ribu"
	case a >= 100000 && a <= 999999:
		spell = "Ratusan ribu"
	default:
		spell = "tak ada input yang sesuai"
	}
	fmt.Println(spell)
}	