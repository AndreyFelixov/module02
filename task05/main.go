package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := make([]string, 0, 5)
	a = append(a, "Hello", "my", "name", "is", "Andrey")
	x := " "
	fmt.Println(Contains(a, x))
	fmt.Println(getMax(3, 4, 52, 9, -232, 1, 3, 5, 6, 7))
}

func Contains(a []string, x string) bool {
	b := strings.Join(a, " ")
	return strings.Contains(b, x)

}
func getMax(c ...int) int {
	sort.Ints(c)
	return c[len(c)-1]
}
