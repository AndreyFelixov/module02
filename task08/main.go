package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var countStr int

func main() {
	pathIn := "data/in.txt"
	c, _ := getFile(pathIn)
	fmt.Print("Total strings:", c)

}

func getFile(pathIn string) (int, error) {
	//flags := log.LstdFlags | log.Lshortfile
	logInfo := log.New(os.Stdout, "INFO:\t", log.Lshortfile)
	fileIn, err := os.Open(pathIn)

	if err != nil {
		log.Fatalln("Unable to open file")
	}

	s := bufio.NewScanner(fileIn)
	var character rune
	character = s.Scan()

	// for s.Scan() {
	// 	if s.Text() == "" {
	// 		logInfo.Println("EOF")
	// 	} else {
	// 		countStr += 1
	// 	}
	// }

	return countStr, s.Err()
}
