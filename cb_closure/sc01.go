package main

import "fmt"

func adder() func(int) int {
	s := 0
	return func(x int) int {
		s += x
		return s
	}
}

func main00() {
	retFunc := adder()
	fmt.Println(retFunc(2))
	fmt.Println(retFunc(3))

}

// 闭包变量竞争
func f1() (i int) {
	i = 10
	defer func() {
		i += 1
	}()
	return 5

}

func main01() {
	foo := f1()
	fmt.Println(foo)
}

// caller-save demo
func f2() int {
	i := 10
	defer func() {
		i += 1
	}()
	return i

}

func main02() {
	fmt.Println(f2()) // 10
}
func f3() (i int) {
	i = 10
	defer func() {
		i += 1
	}()
	return i
}

func main03() {
	fmt.Println(f3()) // 11
}


func main() {
	main02()
	fmt.Println("===")
	main03()
}
