package src

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

const CdataS = "<![CDATA["
const CdataE = "]]>"
const StringFileStart = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
const ErrFileOpen = "fail to open file "
const ErrFileCreate = "fail to create file "
const ErrFileWrite = "error write file "

func CsvToString(csvFile, stringFile string) {
	fmt.Println("CsvToString start...")
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
	//https://groups.google.com/forum/#!topic/golang-china/HmoS5PPXNzM
	reader := csv.NewReader(transform.NewReader(fileCsv, simplifiedchinese.GB18030.NewDecoder()))
	_, _ = fileStr.WriteString(StringFileStart)
	encoder := xml.NewEncoder(fileStr)
	encoder.Indent("", "  ")
	elemRes := CreateStartElement("resources", nil)
	_ = encoder.EncodeToken(elemRes)
	var key, value string
	for {
		s, err := reader.Read()
		if err != nil {
			break
		}
		// csv标题
		if s[0] == "key" {
			continue
		}
		key = s[0]
		value = s[2]
		elemStr := CreateStartElement("string", [][]string{{"name", key}})
		_ = encoder.EncodeToken(elemStr)
		_ = encoder.EncodeToken(xml.CharData(value))
		_ = encoder.EncodeToken(elemStr.End())
	}
	err = encoder.EncodeToken(elemRes.End())
	PrintErrorIfExist(err)
	err = encoder.Flush()
	PrintErrorIfExist(err)
	fmt.Println("CsvToString end...")
}

func StringToCsv(stringFile, csvFile string) {
	fmt.Println("StringToCsv start...")
	array, err := parseFromFile(stringFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveToFile(csvFile, array)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("StringToCsv end...")
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
			switch t.Name.Local {
			case "resources":
				continue
			case "string":
				key = findKeyNameInString(t)
				value = ""
			default:
				value = value + "<" + t.Name.Local + ">"
			}
		case xml.CharData:
			c := string([]byte(t))
			if strings.Contains(c, "<") && !strings.Contains(c, CdataS) {
				c = CdataS + c + CdataE
			}
			value = value + c
		case xml.EndElement:
			if t.Name.Local == "resources" {
				continue
			}
			switch t.Name.Local {
			case "resources":
				continue
			case "string":
				//Go只支持UTF-8编码，而Excel文件是GB2312编码
				key, _ = UTF82GBK(key)
				value, _ = UTF82GBK(value)
				array = append(array, []string{key, value, ""})
			default:
				value = value + "</" + t.Name.Local + ">"
			}
		}
	}
	return array, nil
}

func findKeyNameInString(t xml.StartElement) string {
	var key = ""
	for _, e := range t.Attr {
		if e.Name.Local == "name" {
			key = e.Value
		}
	}
	return key
}

func saveToFile(path string, infos [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(ErrFileCreate + path)
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)

	err = csvWriter.Write([]string{"key", "value", "translate"})
	if err != nil {
		fmt.Println("warning fail create csv Title")
	}
	err = csvWriter.WriteAll(infos)
	csvWriter.Flush()
	if err != nil {
		return errors.New(ErrFileWrite)
	}
	return nil
}
