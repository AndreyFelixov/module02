package main

import (
	"fmt"
	"math"
)

func main() {
	var R *float64
	var P float64
	P = 35
	var r float64

	R = &r

	r = P / 2 * math.Pi
	fmt.Printf("%.2f \n", *R)

	var S float64

	S = math.Pi * (math.Pow(*R, 2))
	fmt.Printf("%.2f \n", S)

}
