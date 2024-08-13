package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "data/06_task_in.txt"
	path1 := "data/out.txt"
	getFile(path, path1)

}

func getFile(path, path1 string) {
	//i := 1
	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)

	f1, _ := os.Open(path1)
	defer f1.Close()

	w := bufio.NewWriter(f1)
	//b := strconv.Itoa(i)

	for s.Scan() {

		// w.WriteString(b)
		// w.WriteString(".")
		// w.WriteString(s.Text())
		// w.WriteString("\n")
		// i += 1
		fmt.Println(s.Text())

	}

	w.Flush()

}
