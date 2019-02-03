package demo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunScan() {
	DemoTitle("console input")
	//test()
	//run()

	read := bufio.NewReader(os.Stdin)
	fmt.Println("input S to end")
	inputS, _ := read.ReadString('\n')
	arraySpace := strings.Split(inputS, " ")
	stringNoRS := strings.Replace(inputS, "\r", "", -1)
	stringNoRS = strings.Replace(stringNoRS, "\n", "", -1)
	numberN := strings.Count(inputS, "\n")
	fmt.Printf("just input: %s, %d size, %d word,%d row",
		inputS, len(stringNoRS), len(arraySpace), numberN)

}
func run() {
	var nrchars, nrwords, nrlines int

	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		// Windows S\n, on Linux it is "S\r\n"
		fmt.Println(input)
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			break
		}
		Counters(input, nrchars, nrwords, nrlines)
	}
}

func Counters(input string, nrchars, nrwords, nrlines int) {
	nrchars += len(input) - 2 // -1 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}

func test() {
	var name, solg string
	fmt.Println("wait for input name ...")
	_, _ = fmt.Scan(&name, &solg)
	fmt.Printf("just input name = %s,slog = %s \n", name, solg)
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("wait for input name ...")
	// 碰到特定的字符串就停止
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
