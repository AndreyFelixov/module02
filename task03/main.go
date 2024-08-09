package main

import (
	"fmt"
)

func main() {
	Week := make([]string, 7)
	Week = []string{"Pon", "Vt", "Sr", "Cht", "Pt", "Sb", "Vs"}
	workDays := make([]string, 5)
	copy(workDays, Week)
	Week = Week[5:]
	fmt.Println(Week)
	fmt.Println(workDays)

}
