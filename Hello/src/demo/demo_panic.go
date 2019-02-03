package demo

import "fmt"

func RunPanic() {
	DemoTitle("panic")
	// defer-panic-recover,直接try-catch不是更棒
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch exception ----->", err)
		}
	}()
	panic("must shut down")
}
