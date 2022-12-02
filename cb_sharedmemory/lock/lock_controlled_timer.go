package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("start")
	go periodic()
	time.Sleep(3 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	fmt.Println("timer cancelled")
	time.Sleep(1 * time.Second)

}

func periodic() {
	for {
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			mu.Unlock()
			return
		}
	}
}

// 结论：在对外部变量进行读写的前后，必须对该共享变量进行加锁保护。
// 虽然在这个代码里面，把有关锁的部分全部去掉，执行结果一致。
// 但这完全是由并发度不够引起的。
// 这就是所谓的go编程哲学：你不知道不这样写会发生什么样的错误，但是你知道安装这个模式去写能工作良好。
