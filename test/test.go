package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	Name    string
	Address string
	City    string
}

var pntr *int

func main() {
	path := "data/in777.txt"
	path1 := "data/out.txt"
	getFile(path, path1)
}
func getFile(path, path1 string) {
	rowsFile, _ := os.ReadFile(path)
	file, _ := os.Create("data/out.txt")
	writer := bufio.NewWriter(file)
	defer file.Close()
	rowLines := strings.Split(string(rowsFile), "\n")
	fmt.Print(len(rowLines))
	for i := 0; i < len(rowLines); i++ {
		rowLine := strings.Split(string(rowLines[i]), "|")
		newRow := Row{Name: rowLine[0], Address: rowLine[1], City: rowLine[2]}
		if newRow.Name != "" && newRow.Address != "" && newRow.City != "" {
			writer.WriteString(strconv.Itoa(i + 1))
			writer.WriteString("\n")
			writer.WriteString(newRow.Name)
			writer.WriteString("\n")
			writer.WriteString(newRow.Address)
			writer.WriteString("\n")
			writer.WriteString(newRow.City)
			writer.WriteString("\n\n\n")
		} else {
			outFile, _ := os.ReadFile(path1)
			fmt.Print(string(outFile))
			pntr = &i
		}
		writer.Flush()
	}
	err := someFunc()
	if err != nil {
		fmt.Println(err)
	}
}
func someFunc() (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "parse error":
				err = fmt.Errorf("parse error: empty field in %dth user", *pntr+1)
			default:
				panic("critical")
			}
		}
	}()
	panic("parse error")
}
