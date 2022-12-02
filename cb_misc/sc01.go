package main

import "fmt"

func visit(ls []int, f func(int)) {
	for _, l := range ls {
		// 执行回调函数
		f(l)
	}
}
func main() {
	visit([]int{1, 2, 3}, func(v int) {
		fmt.Println(v+1)
	})
}
