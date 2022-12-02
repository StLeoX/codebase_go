package main

import (
	"fmt"
	"sync"
)

func demo01() {
	var a string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "hello"
		wg.Done()
	}()
	wg.Wait()
	println(a)
}

///demo02, attention for these different execution sequence
func sendRPC(x int) {
	fmt.Println(x)
}

func demo0201() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			sendRPC(x)
			wg.Done()
		}(i)
		wg.Wait() // acquire lock, then next yield
	}
}

// pass i to anonymous goroutine through argument
func demo0202() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			sendRPC(x)
			wg.Done()
		}(i)
	}
	wg.Wait() //all yield, then acquire lock
}

// pass i to anonymous goroutine via closure
//! this is wrong
func demo0203() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			sendRPC(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	demo0203()
}
