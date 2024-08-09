package main

import (
	"fmt"
)

func main() {
	//Week := make([]string, 0, 7)
	workDays := []string{"Pon", "Vt", "Sr", "Cht", "Pt"}
	weekend := []string{"Sb", "Vs"}

	week := append(workDays, weekend...)

	fmt.Println(week)
}
