package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	pathIn := "data/in.txt"
	getFile(pathIn)
}

func getFile(pathIn string) {
	fileIn, err := os.Open(pathIn)
	if err != nil {
		log.Println("Can't open file: wrong path")
	}
	s := bufio.NewScanner(fileIn)
	var countStr int
	for s.Scan() {
		countStr += 1
	}
	fmt.Printf("Total strings: %d\n", countStr)
	defer fileIn.Close()

}
