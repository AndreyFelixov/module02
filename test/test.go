package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/projects/module02/task08/data/in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//data := make([]byte, 64)
	var str int

	for {
		_, err := os.ReadFile("D:/projects/module02/task08/data/in.txt")
		if err == io.EOF { // если конец файла
			fmt.Println("eof")
			break
		}
		str += 1
	}
	fmt.Println(str)
}
