package demo

import (
	"fmt"
	"time"
)

var cha1Valve, cha2Valve int

func RunGoroutinesSelect() {
	fmt.Println("...RunGoroutinesSelect...")
	ch1, ch2 := make(chan int), make(chan int)
	go runChannel1(ch1)
	go runChannel2(ch2)
	go checkResult(ch1, ch2)
	time.Sleep(1e9)
	fmt.Println("1 s two routines result ", cha1Valve, cha2Valve)
}

func checkResult(ch1 chan int, ch2 chan int) {
	for {
		select {
		case _ = <-ch1:
			cha1Valve++
		case _ = <-ch2:
			cha2Valve++
		}
	}
}

func runChannel2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func runChannel1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
