package main

import (
	"fmt"
	"sync"
)

// using condition variable instead of
// spinlock(infinite loop and check)

func main() {
	// shared
	count := 0
	finished := 0
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			//* When changed shared data, should broadcast cond!
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait() // time.sleep 变为 cond.Wait，自旋变睡眠。
	}
	if count >= 5 {
		fmt.Println("received 5+")
	} else {
		fmt.Println("lost")
	}
	mu.Unlock()
}

/** summary: cond using pattern

mu.Lock()
// do something that might affect your_condition
cond.Broadcast()
mu.Unlock()

---

mu.Lock()
while your_condition == false{
	cond.Wait()
}

// now your_condition is true, so you can do another things when it's true.
mu.Unlock()

*/
