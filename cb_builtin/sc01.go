package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	fmt.Printf("%d", numCPUs)
}
