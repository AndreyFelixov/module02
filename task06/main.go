package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var i int = 1
var b string

func main() {

	t := time.Now()

	p := "data/06_task_in.txt"
	getFile(p)

	logTime := func() float64 {
		return time.Since(t).Seconds()
	}

	fmt.Printf("Использовано байт: %v\n", byteCount())
	fmt.Printf("Выведено строк: %v\n", i-1)
	fmt.Printf("Время выполнения программы:%v\n", logTime())

}

func getFile(path string) {

	sls := make([]string, 0, 100)
	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)

	file, err := os.Create("data/out.txt")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for s.Scan() {

		sls = append(sls, s.Text())

	}

	for _, sl := range sls {

		writer.WriteString(stringCount())
		writer.WriteString(".")
		writer.WriteString(sl)
		writer.WriteString("\n")

	}
	writer.Flush()
}
func byteCount() int {
	var counter int
	file, _ := os.Open("data/out.txt")
	s := bufio.NewScanner(file)
	for s.Scan() {

		counter += len(s.Text())

	}
	return counter

}
func stringCount() string {
	b = strconv.Itoa(i)
	i += 1
	return b
}
