package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	fmt.Println("Channels")

	// Will not run as Channel blocks
	// ch := make(chan int)
	// ch <- 42
	// fmt.Println(<-ch)

	// 1- pass of value
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)

	// 2- buffer usage
	ch = make(chan int, 2)
	ch <- 42
	ch <- 43
	//will cause a blocking
	//ch <- 44
	fmt.Println(<-ch, <-ch)

	fmt.Println("-------------------------------------")

	fmt.Println("Directional Channels")

	ch = make(chan int)
	chs := make(chan<- int)
	chr := make(<-chan int)

	fmt.Printf("ch\t%T\n", ch)
	fmt.Printf("chs\t%T\n", chs)
	fmt.Printf("chr\t%T\n", chr)

	// invalid operation to send to receive-only type <-chan int
	// chr <- 42
	// valid
	// <-chr

	//valid
	// chs <- 42
	// invalid operation receive from send-only type chan<- int
	// <-chs

	// specific to general doesn't assign
	// ch = chr
	// ch = chs
	// general to specific assign
	chr = ch
	chs = ch

	fmt.Println("-------------------------------------")

	fmt.Println("Using Channels")

	go foo(ch)
	bar(ch)
	fmt.Println("Code exited")

	fmt.Println("-------------------------------------")

	fmt.Println("Range")

	// Range stops reading from a channel when the channel is closed
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("value from channel", v)
	}
	fmt.Println("closing range block")

	fmt.Println("-------------------------------------")

	fmt.Println("Select")

	o := make(chan int)
	e := make(chan int)
	f := make(chan int)

	// send
	go send(o, e, f)

	// recieve
	receive(o, e, f)
	fmt.Println("closing select block")

	fmt.Println("-------------------------------------")

	fmt.Println("Comma ok idiom")

	ch = make(chan int)
	go func() {
		ch <- 42
		close(ch)
	}()

	v, ok := <-ch
	fmt.Println("ok idiom value; v:", v, "ok:", ok)

	v, ok = <-ch
	fmt.Println("ok idiom value; v:", v, "ok:", ok)

	fmt.Println("-------------------------------------")

	fmt.Println("Interesting code")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	fanin := make(chan int)

	go sendfun(even, odd, quit)

	go receivefun(even, odd, quit, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

	fmt.Println("-------------------------------------")

	fmt.Println("Fan in")

	val := fanIn(boring("Ajay"), boring("Raj"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-val)
	}
	fmt.Println("You both are boring, I'm leaving")

	fmt.Println("-------------------------------------")

	fmt.Println("Fan out")

	c1 := make(chan int)
	c2 := make(chan int)
	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("All heavy weight opeartion done")

	fmt.Println("-------------------------------------")

	fmt.Println("Context")

	cx := context.Background()

	fmt.Println("context:\t", cx)
	fmt.Println("context err:\t", cx.Err())
	fmt.Printf("context type:\t%T\n", cx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(cx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	fmt.Println("-------------------------------------")
}

// Channels ------------------------------
func foo(ch chan<- int) {
	ch <- 42
}

func bar(ch <-chan int) {
	fmt.Println(<-ch)
}

// Select ---------------------------------
func send(o, e, f chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	f <- 0
}

func receive(o, e, f <-chan int) {
	for {
		select {
		case v := <-o:
			fmt.Println("recieve odd value", v)
		case v := <-e:
			fmt.Println("recieve eve value", v)
		case v := <-f:
			fmt.Println("recieve f value", v)
			return
		}
	}
}

// Interesting code ------------------------------
// send channel
func sendfun(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receivefun(even, odd, quit <-chan int, fanin chan<- int) {
	// var wg sync.WaitGroup // required as v is the value passed between threads
thisLoop:
	for {
		select {
		case v := <-even:
			fanin <- v
			// wg.Add(1)
			// func(v int) {
			// 	fanin <- v
			// wg.Done()
			// }(v)
		case v := <-odd:
			fanin <- v
			// wg.Add(1)
			// func(v int) {
			// 	fanin <- v
			// wg.Done()
			// }(v)
		case <-quit:
			break thisLoop
		}
	}
	// wg.Wait()
	close(fanin)
}

// Fan In ------------------
func boring(name string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%v %v", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(f, s <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-f
		}
	}()
	go func() {
		for {
			ch <- <-s
		}
	}()
	return ch
}

// Fan Out ---------------------
func populate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- heavyOperation(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func heavyOperation(v int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1e3)
}
