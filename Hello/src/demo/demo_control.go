package demo

import (
	"fmt"
	"strconv"
)

func RunControl() {
	DemoTitle("Control")
	labelContinueRun()
}

func labelContinueRun() {
	//标签的作用对象为外部循环，因此 i 会直接变成下一个循环的值，而此时 j 的值就被重设为 0，即它的初始值.所以没有 2，3 i的输出
	/*
	run i = 0; j = 0
	run i = 0; j = 1
	run i = 1; j = 0
	run i = 1; j = 1
	 */

LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if 2 == j {
				continue LABEL1
				//continue
			}
			fmt.Println("run i = " + strconv.Itoa(i) + "; j = " + strconv.Itoa(j))
		}
	}
}
