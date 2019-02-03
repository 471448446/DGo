package main

import "fmt"

func main() {
	//以下会永远阻塞
	ch := make(chan int)
	//ch <- 2
	fmt.Println(<-ch)
}
