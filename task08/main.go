package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	pathIn := "data/in.txt"
	c, _ := getFile(pathIn)
	fmt.Print("Total strings: ", c)

}

func getFile(pathIn string) (int, error) {
	logInfo := log.New(os.Stdout, "INFO:\t", log.Lshortfile)
	fileIn, err := os.Open(pathIn)
	var countStr int
	defer func() {
		if fileIn != nil {
			fileIn.Close()
		} else {
			log.Fatalln("File doesn't exist")
		}
	}()

	if err != nil {
		log.Fatalln("Unable to open file")
	}
	r := bufio.NewReader(fileIn)
	for {
		_, err = r.ReadString('\n')
		countStr += 1
		if err != nil {
			break
		}
	}
	if err != io.EOF {

		return 0, err
	} else {
		logInfo.Println("eof")
	}

	return countStr, err
}
