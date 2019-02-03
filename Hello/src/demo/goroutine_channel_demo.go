package demo

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var ErrTimeout = errors.New("TimeOutException")
var ErrInterrupt = errors.New("InterruptException")

type Runner struct {
	//任务列表，一堆可执行的函数
	tasks     [] func(int)
	complete  chan error
	interrupt chan os.Signal
	/*
	倒计时，只能接收的通道，这里的话倒计时完了会收到一个通道消息
	在Start()方法里面的select 如果收到了timeOut消息就会返回一个ErrTimeout的error
	 */
	timeOut <-chan time.Time
}

func NewRunner() Runner {
	return Runner{
		complete: make(chan error),
		//我们可以至少接收到一个操作系统的中断信息，
		// 这样Go runtime在发送这个信号的时候不会被阻塞，如果是无缓冲的通道就会阻塞了
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) run() error {
	for key, task := range r.tasks {
		if r.Interrupt() {
			return ErrInterrupt
		}
		// 耗时key 秒
		task(key)
	}
	return nil
}

func (r *Runner) Interrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Add(task func(int)) {
	r.tasks = append(r.tasks, task)
}

func (r *Runner) Adds(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) Start(tm time.Duration) error {
	r.timeOut = time.After(tm)
	//系统中断的信号，发给r.interrupt即可
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		// 发送结果数据
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeOut:
		return ErrTimeout
	}
}

func RunGoroutinesDemoRunner() {
	fmt.Println("RunnerTask start...")
	runner := NewRunner()

	runner.Adds(createTask(), createTask(), createTask())

	if err := runner.Start(4e9); err != nil {
		fmt.Println(err)
		switch err {
		case ErrTimeout:
			os.Exit(-1)
		case ErrInterrupt:
			os.Exit(-2)
		}
	}
	fmt.Println("RunnerTask finish...")
}

func createTask() func(int) {
	return func(i int) {
		fmt.Println("run task ", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
