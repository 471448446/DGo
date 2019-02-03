package main

import "fmt"
import "./src/demo"

func main() {
	fmt.Println("Hello Go", demo.Version)
	demo.RunConst()
	demo.RunString()
	demo.RunTime()
	demo.RunControl()
	demo.RunClose()
	//demo.RunScan()
	demo.RunFileRead()
	demo.RunPanic()
	demo.RunGoroutines()
}
