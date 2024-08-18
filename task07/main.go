package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello world")
	path := "data/07_task_in.txt"
	getFile(path)
}

func getFile(path string) {
	file, _ := os.Open(path)
	defer file.Close()

}
