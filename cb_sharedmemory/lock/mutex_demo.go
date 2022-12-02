package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i_ int) {
			fmt.Printf("start %d\n", i_)
			lock.RLock()
			fmt.Printf("read %d\n", i_)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	fmt.Println("准备释放写锁")
	time.Sleep(time.Second * 2)
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()

}
