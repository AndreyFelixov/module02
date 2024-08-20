package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var c int = 1
var b string

type Row struct {
	Name    string
	Address string
	City    string
}

func main() {
	path := "data/07_task_in.txt"
	getFile(path)
	fmt.Println("zaebalo")

}

func getFile(path string) {
	rows := make([]Row, 0)
	rowsFile, _ := os.ReadFile(path)
	rowLines := strings.Split(string(rowsFile), "\n")
	for i := 0; i < len(rowLines); i++ {
		if rowLines[i] != "" {
			rowLine := strings.Split(string(rowLines[i]), "|")
			newRow := Row{Name: rowLine[0], Address: rowLine[1], City: rowLine[2]}
			rows = append(rows, newRow)
		} else {
			panic("parse error: empty field on string")
		}
	}

	file, _ := os.Create("data/out.txt")
	writer := bufio.NewWriter(file)
	defer file.Close()
	for _, roww := range rows {
		if roww.Name != "" && roww.Address != "" && roww.City != "" {
			writer.WriteString(stringCount())
			writer.WriteString("\n")
			writer.WriteString(roww.Name)
			writer.WriteString("\n")
			writer.WriteString(roww.Address)
			writer.WriteString("\n")
			writer.WriteString(roww.City)
			writer.WriteString("\n\n\n")
		} else {

			writer.Flush()
			fmt.Printf("parse error:empty field on string %s\n", stringCount())
			panic("")
		}
	}
	writer.Flush()
}

func stringCount() string {
	b = strconv.Itoa(c)
	c += 1
	return b
}
