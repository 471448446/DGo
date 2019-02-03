package demo

import "fmt"

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
	Sunday                     = 7
)

//使用iota来简写这种增长的常量
//遇到const，iota的值变为0
const (
	Monday2 = iota + 1
	Tuesday2
	Wednesday2
	Thursday2
	Friday2
	Saturday2
	Sunday2
)
const name = iota

func RunConst() {
	DemoTitle("Const")
	fmt.Printf("Monday2 = %d, Tuesday2=%d, Wednesday2= %d, Thursday2= %d, "+
		"Friday2 = %d,Staturday2= %d,Sunday2 = %d \n",
		Monday2, Tuesday2, Wednesday2, Thursday2, Friday2, Saturday2, Sunday2)
	fmt.Printf("iota = %d", name)
}
