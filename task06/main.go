package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func main() {

	t := time.Now()

	p := "data/06_task_in.txt"
	getFile(p)

	logTime := func() float64 {
		//t1 :=
		return time.Since(t).Seconds()
	}
	//fmt.Printf("Время выполнения программы:%v\n", logTime)

	fmt.Printf("Время выполнения программы:%v\n", logTime())

}

func getFile(path string) {
	i := 1
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
	defer func() {
		fmt.Printf("Выведено строк: %v\n", i-1)
		fmt.Printf("Использовано байт:%v\n", unsafe.Sizeof(sls))
		file.Close()
	}()

	for s.Scan() {

		sls = append(sls, s.Text())

	}
	for _, sl := range sls {

		b := strconv.Itoa(i)
		writer.WriteString(b)
		writer.WriteString(".")
		writer.WriteString(sl)
		writer.WriteString("\n")
		i += 1
	}
	writer.Flush()
}
