package main

import (
	"fmt"
	"time"
)

func periodic() {
	for {
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("start")
	go periodic()
	time.Sleep(5 * time.Second) // main goroutine is sleeping, to keep this program not exit.
}
