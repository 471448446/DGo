package demo

import (
	"fmt"
	"time"
)

func RunGoroutines() {
	DemoTitle("Goroutines")
	//tryBlock()
	routineOne()
	runTestBlock3()
	RunGoroutinesSync()
	runTest()
	RunGoroutinesDemoRunner()
	RunGoroutinesSelect()

}

func tryBlock() {
	//以下报错fatal error: all goroutines are asleep - deadlock!
	//实际上却没有
	ch := make(chan int)
	// 直接发
	ch <- 2
	// 直接收
	fmt.Println(<-ch)
}

/*
模拟交互
 */
func runTest() {
	fmt.Println("...client request server...")
	ch := make(chan string)

	go func() {
		fmt.Println("server: server handle url...")
		//做耗时操作
		//time.Sleep(1e9)
		ch <- "hello"
		//由于只接受了一次消息，所以这句话是接受不到的
		ch <- "dear client"
	}()
	fmt.Println("client: request to server")
	/*
	等待输出的时候挂起，下面一行就不会执行
	这个时候才会去执行另外一个协程
	 */
	fmt.Println("client: response from server : ", <-ch)
	fmt.Println("client: talk to server finish")
}

/*
写一个通道证明它的阻塞性，开启一个协程接收通道的数据，持续 15 秒，然后给通道放入一个值。
 */
func runTestBlock3() {
	fmt.Println("...runTestBlock3...")
	ch := make(chan int)
	go func() {
		time.Sleep(5e9)
		x := <-ch
		fmt.Println("after 5s ,get data", x)
	}()
	fmt.Println("prepare send data", 10)
	ch <- 10
	//这句是在接收到数据后打印出来的
	fmt.Println("send data", 10)
}

func routineOne() {
	fmt.Println("...run 1s...")
	/*
	main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
	getData() 使用了无限循环：它随着 sendData() 的发送完成和 ch 变空也结束了。
	如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic
	 */
	chOne := make(chan string)
	// 启动两个协程
	go sendData(chOne)
	go getData(chOne)
	/*
	1e9 = 1 s? Sleep 的单位是纳秒
	给这次运行1秒,Sleep会阻塞当前的协程，这里就是主协程，
	主协程被阻塞的时候，调度器回去唤醒其他的协程运行，这里就是执行sendData和getData
	 */
	time.Sleep(1 * time.Second)
}

func getData(ch chan string) {
	fmt.Println(".....getData start")
	var input string
	/*
	如果开启这行，是得不到结果的，
	因为第一个协程发送了数据，但是这个协程被睡眠了，接受不到数据
	 */
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Println(input)
	}
}

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	ch <- "Today"
	ch <- "is"
	ch <- "Tuesday"
	ch <- ".....sendData over"
}
