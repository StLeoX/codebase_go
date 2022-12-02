package main

import (
	"fmt"
	"time"
)

// 信道写端
type Sender = chan<- int

// 信道读端
type Recver = <-chan int

func demo0201() {

	var fifo = make(chan int, 10)
	go func() {
		var sender Sender = fifo
		fmt.Println("sending:100")
		sender <- 100
	}()
	go func() {
		var recver Recver = fifo
		recvData := <-recver
		fmt.Printf("recved %d", recvData)
	}()
	time.Sleep(time.Second)
}

// 验证chan的阻塞性
func demo0202() {
	ch := make(chan int, 0)
	go func() {
		fmt.Println("before send")
		ch <- 1
		ch <- 1
		fmt.Println("after send")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("preparing recv")
		msg := <-ch
		fmt.Println("received: ", msg)
	}()
	time.Sleep(2 * time.Second)
}

// 验证chan的读端必须先准备好写端才能写
func demo0203() {
	ch := make(chan int, 0)
	go func() {
		time.Sleep(1 * time.Second)
		<-ch
	}()
	start := time.Now()
	ch <- 1
	fmt.Printf("send took %v\n", time.Since(start))
	/** 切忌写出这样的线性代码：
	<-ch
	ch<-1
	*/

}

func main() {
	ch:=make(chan int,0)
	ch<-1
	<-ch
	fmt.Println("fine")
}
