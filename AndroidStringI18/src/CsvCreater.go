package src

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

func CsvToString(csvFile, stringFile string) {
	fileCsv, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(ErrFileOpen + csvFile)
		return
	}
	fileStr, err := os.Create(stringFile)
	if err != nil {
		fmt.Println(ErrFileCreate + csvFile)
		return
	}
	defer fileCsv.Close()
	defer fileStr.Close()
	reader := csv.NewReader(fileCsv)
	encoder := xml.NewEncoder(fileStr)
	encoder.Indent("", "  ")

	elemRes := StartElement("resources", nil)
	_ = encoder.EncodeToken(elemRes)
	var key, value string
	for {
		s, err := reader.Read()
		if err != nil {
			break
		}
		if s[0] == "key" {
			continue
		}
		fmt.Println(value)
		elemStr := StartElement("string", [][]string{{"name", key}})
		_ = encoder.EncodeToken(elemStr)
		_ = encoder.EncodeToken(xml.CharData(value))
		_ = encoder.EncodeToken(elemStr.End())
	}
	_ = encoder.EncodeToken(elemRes.End())
}

func StringToCsv(stringFile, csvFile string) {
	fmt.Println("start...")
	array, err := parseFromFile(stringFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveToFile(csvFile, array)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("end...")
	}
}

func parseFromFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(ErrFileOpen + path)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	p := xml.NewDecoder(reader)
	key, value := "", ""
	var array [][]string
	for token, err := p.Token(); err == nil; token, err = p.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "resources" {
				continue
			}
			for _, e := range t.Attr {
				if e.Name.Local == "name" {
					key = e.Value
				}
			}
		case xml.CharData:
			value = string([]byte(t))
		case xml.EndElement:
			if t.Name.Local == "resources" {
				continue
			}
			//Go只支持UTF-8编码，而Excel文件是GB2312编码
			key, _ = UTF82GBK(key)
			value, _ = UTF82GBK(value)
			array = append(array, []string{key, value, ""})
		}
	}
	return array, nil
}

func saveToFile(path string, infos [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(ErrCsvCreate + path)
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)

	err = csvWriter.Write([]string{"key", "value", "translate"})
	if err != nil {
		fmt.Println("warning to create csv Title")
	}
	err = csvWriter.WriteAll(infos)
	csvWriter.Flush()
	if err != nil {
		return errors.New(ErrCsvWrite)
	}
	return nil
}
