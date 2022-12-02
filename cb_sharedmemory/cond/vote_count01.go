package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	finished := 0
	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			if vote {
				count++
			}
			finished++
		}()
	}
	for count < 5 && finished != 10 {
		//wait
		// 优化方式一：自旋锁（spinlock）
		// 自旋锁就是这样，在无限循环中不断加锁、释放，间歇地进行检查。
		time.Sleep(50 * time.Millisecond)
		// 那么新的问题是，50ms或者是多少时间相对于检查操作是合适的。

	}
	if count >= 5 {
		fmt.Println("received 5+")
	} else {
		fmt.Println("lost")
	}

}

func requestVote() bool {
	a := rand.Intn(2)
	if a == 0 {
		return false
	} else {
		return true
	}
}
