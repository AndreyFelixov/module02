package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rows = make([]string, 0)
var pntr *int

func main() {
	pathIn := "D:/projects/module02/task07/data/in.txt"
	pathOut := "D:/projects/module02/task07/data/out.txt"
	getFile(pathIn, pathOut)
}

func getFile(pathIn, pathOut string) {
	flag := false
	fileIn, _ := os.Open(pathIn)
	s := bufio.NewScanner(fileIn)
	var strCounter int
	for s.Scan() {
		strCounter += 1
		if !(strings.Contains(s.Text(), "||") || string(s.Text()[0]) == "|" || string(s.Text()[len(s.Text())-1]) == "|") {
			rows = append(rows, s.Text())
		} else {
			flag = true
			pntr = &strCounter
			break
		}
	}
	fileIn.Close()
	if !flag {
		writeAndPrint(pathOut)
	} else {
		writeAndPrint(pathOut)
		err := someFunc()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func someFunc() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "parse error":
				err = fmt.Errorf("parse error: empty field in %dth string", *pntr)
			default:
				panic("critical")
			}
		}
	}()
	panic("parse error")
}

func writeAndPrint(pathOut string) {
	fileOut, _ := os.Create(pathOut)
	w := bufio.NewWriter(fileOut)
	for i := 0; i < len(rows); i++ {
		rowLine := strings.Split(rows[i], "|")
		w.WriteString(strconv.Itoa(i + 1))
		w.WriteString("\n")
		w.WriteString(rowLine[0])
		w.WriteString("\n")
		w.WriteString(rowLine[1])
		w.WriteString("\n")
		w.WriteString(rowLine[2])
		w.WriteString("\n\n\n")
	}
	w.Flush()
	fileOut, _ = os.Open(pathOut)
	defer fileOut.Close()
	s1 := bufio.NewScanner(fileOut)
	for s1.Scan() {
		fmt.Println(s1.Text())
	}
}