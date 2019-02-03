package demo

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	salary int32
	wg     sync.WaitGroup
	metu   sync.Mutex
)

func RunGoroutinesSync() {
	fmt.Println("...RunGoroutinesSync...")
	wg.Add(2)
	// result 2?
	//go addSalary()
	//go addSalary()
	// result 2?
	go addSalaryAtom("A")
	go addSalaryAtom("B")
	// result 4
	//go addSalaryMutex("A")
	//go addSalaryMutex("B")

	//go takeMoney()
	//go saveMoney()
	//go takeMoneyUseAtomic()
	//go saveMoneyUseAtomic()
	wg.Wait()
	fmt.Println("my money:", salary)
}

func addSalary() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		m := salary
		runtime.Gosched()
		m++
		salary = m
	}
}

//原子锁并没有锁住，结果始终时2
func addSalaryAtom(tag string) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		m := atomic.LoadInt32(&salary)
		fmt.Println("Atom", tag, "in", m)
		runtime.Gosched()
		fmt.Println("Atom", tag, "add", m)
		m++
		atomic.StoreInt32(&salary, m)
		fmt.Println("Atom", tag, "out", m)
	}
}
func addSalaryMutex(tag string) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		metu.Lock()
		m := atomic.LoadInt32(&salary)
		//fmt.Println("Mutex", tag, "in", m)
		runtime.Gosched()
		m++
		//fmt.Println("Mutex", tag, "add", m)
		atomic.StoreInt32(&salary, m)
		//fmt.Println("Mutex", tag, "out", m)
		metu.Unlock()
	}
}

func saveMoney() {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		salary ++
		fmt.Println("salary save:", salary)
		runtime.Gosched()
	}
}

func saveMoneyUseAtomic() {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		m := atomic.LoadInt32(&salary) + 1
		atomic.StoreInt32(&salary, m)
		fmt.Println("salary save:", salary)
		runtime.Gosched()
	}
}

func takeMoney() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		if salary > 1 {
			salary -= 2
			fmt.Println("salary take least:", salary)
		}
		runtime.Gosched()
	}
}

func takeMoneyUseAtomic() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		m := atomic.LoadInt32(&salary)
		if m > 1 {
			atomic.StoreInt32(&salary, m-2)
			fmt.Println("salary take least:", salary)
		}
		runtime.Gosched()
	}
}
