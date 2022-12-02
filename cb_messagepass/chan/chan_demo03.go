package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel usecase
// main is data-receiver
func demo0301() {
	ch := make(chan int)
	for i := 0; i < 4; i++ {
		go doWork(ch)
	}
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
func doWork(ch chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		ch <- rand.Int()
	}
}

// no-buffer channel usecase
// using channel as mutex
func demo0302() {
	done := make(chan bool)
	for i := 0; i < 4; i++ {
		go func(x int) {
			sendRPC(x)
			done <- true
		}(i)
	}
	for i := 0; i < 4; i++ {
		<-done // ReadEnd controlling the data generation
	}

}

func sendRPC(x int) {
	fmt.Println(x)
}

func main() {
	demo0302()
}
