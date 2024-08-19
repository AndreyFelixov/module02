package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var i int = 1
var b string

type Row struct {
	Name    string
	Address string
	City    string
}

func main() {
	path := "data/07_task_in.txt"
	getFile(path)

}

func getFile(path string) {
	sls := make([]string, 0)
	f, _ := os.Open(path)
	defer f.Close()
	s := bufio.NewScanner(f)
	file, err := os.Create("data/out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for s.Scan() {

		sls = append(sls, s.Text())

	}
	rowLines := strings.Split(string(sls[1]), "\n")

}
func stringCount() string {
	b = strconv.Itoa(i)
	i += 1
	return b
}
