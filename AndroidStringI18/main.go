package main

import (
	"./src"
	"fmt"
	"os"
	"strings"
)

func main() {
	//  go run D:/WP/gitHub/DGo/AndroidStringI18/main.go D:/strings.xml D:/strings.csv
	file1, file2 := os.Args[1], os.Args[2]

	if file1 == "" || file2 == "" {
		fmt.Printf("warning files f1 = %s,f2 = %s\n", file1, file2)
		file1 = "D:/strings.xml"
		file2 = "D:/strings.csv"
		//file1 = "D:/strings.csv"
		//file2 = "D:/stringsT.xml"
	}

	suffixCsv, suffixXml := ".csv", ".xml"
	if strings.Contains(file1, suffixXml) && strings.Contains(file2, suffixCsv) {
		src.StringToCsv(file1, file2)
	} else if strings.Contains(file1, suffixCsv) && strings.Contains(file2, suffixXml) {
		src.CsvToString(file1, file2)
	} else {
		fmt.Println("error not support parse way")
	}
	fmt.Println("file saved in ", file2)
}
