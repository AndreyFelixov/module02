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

func main() {
	path := "data/in.txt"
	path1 := "data/out.txt"
	getFile(path, path1)
}
func getFile(path, path1 string) {
	rowsFile, _ := os.ReadFile(path)
	file, _ := os.Create("data/out.txt")
	writer := bufio.NewWriter(file)
	defer file.Close()
	rowLines := strings.Split(string(rowsFile), "\n")
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
			writer.WriteString("\n\n\n\n")
		} else {
			outFile, _ := os.ReadFile(path1)
			fmt.Print(string(outFile))
			fmt.Printf("parse error:empty field on string %d\n", i)
			panic("")
		}
		writer.Flush()
	}
}
