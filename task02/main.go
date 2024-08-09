package main

import (
	"fmt"
)

func main() {

	var A *int
	B := 322
	A = &B

	fmt.Println(*A)

	*A = 200
	B = *A
	fmt.Println(B)

}
