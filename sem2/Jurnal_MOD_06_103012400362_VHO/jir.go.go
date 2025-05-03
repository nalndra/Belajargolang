package main
import "fmt"

const bigParcelPaper int = 60 * 80
const smallParcelPaper int = 40 * 60
const paperSize int = 80 * 100
const bigParcelCost int = 75000
const smallParcelCost int = 45000
const paperCost int = 750

func main(){
	var employee string
	var highLevel, entryLevel, totalEmployeesVal int
	var totalPaperNeededSizeVal, totalPaperNeededVal, totalCostVal int

	for employee != "#"{
		checkEmployees(employee, &highLevel, &entryLevel)
		fmt.Scanln(&employee)
	}
	totalEmployeesVal = totalEmployees(highLevel, entryLevel)
	totalPaperNeededSizeVal = totalPaperNeededSize(highLevel, entryLevel)
	totalPaperNeededVal =  totalPaperNeeded(totalPaperNeededSizeVal)
	totalEmployeesVal = (totalEmployees(highLevel, entryLevel))
	updateTotalPaperNeeded(&totalPaperNeededVal, needMore(totalPaperNeededSizeVal))
	totalCostVal = totalCost(totalPaperNeededVal, highLevel, entryLevel)
	
	fmt.Printf("Vcorp membutuhukan anggaran sebesar Rp %d, untuk\n", totalCostVal)
	fmt.Printf("memberikan parcel kepada %d pegawai.", totalEmployeesVal)
	fmt.Printf("Yang terdiri dari %d bigParcel dan %d smallParcel", highLevel, entryLevel)
}

func checkEmployees(position string,  highLevel, entryLevel *int) {
	if position == "manager" || position == "supervisor"{
		*highLevel++
	}else if position == "officer" || position == "staff" || position == "intern"{
		*entrylevel++
	}
}

func totalEmployees(highLevel, entryLevel int) int{
	return highLevel + entryLevel
}

func totalPaperNeededSize(highLevel, entryLevel int) int{
	return (highLevel * bigParcelPaper) + (entryLevel * smallParcelPaper)
}

func totalPaperNeeded(totalPaperNeeded int) int{
	return (totalPaperNeededSize / paperSize)
}

func needMore(totalPaperNeededSize int) bool{
	if totalPaperNeededSize % paperSize != 0{
		return true
	}
	return false
}

func updateTotalPaperNeeded(totalPaperNeeded *int, needMore bool){
	if needMore == true{
		*totalPaperNeeded++
	}
}

func totalCost(totalPaperNeeded, highLevel, entryLevel int) int{
	var parcelCost int
	var paperCosts int
	
	paperCosts = totalPaperNeeded * paperCost
	parcelCost = (highLevel * bigParcelCost) + (entryLevel * smallParcelCost)
	return paperCosts + parcelCost
}