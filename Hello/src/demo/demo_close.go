package demo

import "fmt"

func RunClose() {
	DemoTitle("close")

	onePlus100 := func() int {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		return sum
	}()
	fmt.Println("1 + ....+100 = ", onePlus100)
	// 返回是2，defer 加了一次
	fmt.Println("defer result ", f())
	fmt.Println("add() ", add()())

	p1, p2 := adder(), adder()

	for i := 0; i < 3; i++ {
		fmt.Printf("p1 %d = %d,p2 %d = %d \n", i, p1(i), i, p2(-i))
	}
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 1
}

func add() func() int {
	sum := 0
	return func() int {
		sum ++
		return sum
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
