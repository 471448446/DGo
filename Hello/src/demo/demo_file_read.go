package demo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func RunFileRead() {
	DemoTitle("ReadFile")
	readCode()
	readInStringAndCopy()
}

func readInStringAndCopy() {
	path := "./data/data.txt"
	pathO := "./data/dataCopy.txt"
	str, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(path, "open fail ", err.Error())
		return
	}
	fmt.Println("read str :", string(str))
	writeWrr := ioutil.WriteFile(pathO, str, 0644)
	if writeWrr != nil {
		fmt.Println("write fail ", writeWrr.Error())
	} else {
		fmt.Println("write ok", pathO)
	}
}

func readCode() {
	path := "./data/hello.txt"
	inputFile, openErr := os.Open(path)
	if openErr != nil {
		fmt.Println("faile open file")
		return
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)
	for {
		str, readErr := reader.ReadString('\n')
		fmt.Print(str)
		if readErr == io.EOF {
			break
		}
	}
	fmt.Printf("\nread %s end \n", path)
}
