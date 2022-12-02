package main

import (
	"fmt"
	"sync"
	"time"
)

// this is a wrong demo
// two mutexes separate single atomic trans: alice++ && bob--.
func demo01() {
	alice := 1000
	bob := 1000
	var mu_alice sync.Mutex
	var mu_bob sync.Mutex

	total := alice + bob

	go func() {
		for i := 0; i < 100; i++ {
			mu_alice.Lock()
			alice += 1
			mu_alice.Unlock()

			mu_bob.Lock()
			bob -= 1
			mu_bob.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			mu_bob.Lock()
			bob += 1
			mu_bob.Unlock()

			mu_alice.Lock()
			alice -= 1
			mu_alice.Unlock()

		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		// multi lock push
		//mu_alice.Lock()
		//mu_bob.Lock()

		if alice+bob != total {
			fmt.Printf("alice:%d,bob:%d,sum:%d\n", alice, bob, alice+bob)
		}

		// multi lock pop
		//mu_bob.Unlock()
		//mu_alice.Unlock()

	}
}

func demo02() {
	alice := 1000
	bob := 1000
	var mu sync.Mutex
	total := alice + bob
	go func() {
		for i := 0; i < 100; i++ {
			mu.Lock()
			alice += 1
			bob -= 1
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			mu.Lock()
			alice -= 1
			bob += 1
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("alice:%d,bob:%d,sum:%d\n", alice, bob, alice+bob)
		}
		mu.Unlock()
	}
}

func main() {
	demo02()
}
