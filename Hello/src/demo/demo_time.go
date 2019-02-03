package demo

import (
	"fmt"
	"time"
)

func RunTime() {
	DemoTitle("time")
	now := time.Now()
	fmt.Println(now.String())
	fmt.Println(now.Unix())
	fmt.Printf("%4d/%02d/%02d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),now.Hour(),now.Minute(),now.Second())
}
