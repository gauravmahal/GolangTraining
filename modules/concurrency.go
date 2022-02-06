package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("WaitGroup")

	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("-------------------")

	fmt.Println("Race condition and then Mutex usage")
	{
		counter := 0
		gs := 100
		wg.Add(gs)
		fmt.Println("CPU\t\t", runtime.NumCPU())
		fmt.Println("Goroutine\t", runtime.NumGoroutine())

		var mx sync.Mutex
		for i := 0; i < gs; i++ {
			go func() {
				mx.Lock()
				v := counter
				runtime.Gosched()
				v++
				counter = v
				mx.Unlock()
				wg.Done()
			}()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}

		wg.Wait()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("counter:", counter)
	}
	fmt.Println("-------------------")

	fmt.Println("Atomic usage")
	{
		var counter int64
		gs := 100
		wg.Add(gs)
		fmt.Println("CPU\t\t", runtime.NumCPU())
		fmt.Println("Goroutine\t", runtime.NumGoroutine())

		for i := 0; i < gs; i++ {
			go func() {
				atomic.AddInt64(&counter, 1)
				fmt.Println("counter:", atomic.LoadInt64(&counter))
				runtime.Gosched()
				wg.Done()
			}()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}

		wg.Wait()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("counter:", counter)
	}
	fmt.Println("-------------------")
}

// WaitGroup ----------------
func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
		runtime.Gosched()
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
		runtime.Gosched()
	}
	wg.Done()
}

//
