package main

import (
	"fmt"
	"sync"
	"time"
)

func unctrl_counter() {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() { counter += 1 }()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)

}
func ctrled_counter() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter += 1
			//mu.Unlock()
			// 不允许加锁一次、释放两次。
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}
func main() {
	unctrl_counter()
	ctrled_counter()
}
