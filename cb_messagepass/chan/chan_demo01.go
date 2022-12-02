package main

import (
	"fmt"
	"time"
)

// 使用 chan bool 实现互斥锁
func atomIncrease(ch chan bool, x *int) {
	ch <- true // lock
	*x = *x + 1
	<-ch //unlock
}

func testAtom() {
	// 注意需要手动设置chan cap
	fifo := make(chan bool, 1)
	x := 0
	for i := 0; i < 100; i++ {
		go atomIncrease(fifo, &x)
	}
	// 执行足够时间
	time.Sleep(time.Second)
	fmt.Println(x)
}

func increase(x *int) {
	*x = *x + 1
}
func testNoAtom() {
	x := 0
	for i := 0; i < 100; i++ {
		go increase(&x)

	}
	fmt.Println(x)
}
func main() {
	testNoAtom()
	testAtom()

}
