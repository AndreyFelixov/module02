package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "104"
	b := 35

	c, _ := strconv.Atoi(a)
	fmt.Println(c)

	v := strconv.Itoa(b)
	fmt.Println(v)
}
