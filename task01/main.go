package main

import "fmt"

func MsToKmh(s float64) float64 {
	return (s * 3600 / 1000)
}

func MsToMh(s float64) float64 {
	return (s * 2.237)
}

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {

	var a EuropeanVelocity

	var s float64
	fmt.Println(a == EuropeanVelocity(s))
	s = 120.4
	a = EuropeanVelocity(MsToKmh(s))
	fmt.Println(a)

	var Am AmericanVelocity
	var b float64
	fmt.Println(Am == AmericanVelocity(b))
	b = 130
	Am = AmericanVelocity(MsToMh(b))
	fmt.Println(Am)

}
