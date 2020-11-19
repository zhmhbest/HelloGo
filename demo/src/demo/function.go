package demo

import "fmt"

//定义函数类型

func Anonymous() {
	var f = func(x int, y int) int {
		return x + y
	}
	fmt.Printf("Anonymous: %d\n", f(1, 2))
	fmt.Print("\n")
}

func MultipleReturns() {
	var f = func() (int, int) {
		return 1, 2
	}
	var i1, i2 = f()
	fmt.Printf("MultipleReturns: (%d, %d)\n", i1, i2)
	fmt.Print("\n")
}

func Closer() {
	type funVar func(int) int
	var getCloser = func() funVar {
		var count = 0
		f := func(x int) int {
			count++
			return count * x
		}
		return f
	}

	var f = getCloser()
	fmt.Printf("Closer: %d\n", f(2))
	fmt.Printf("Closer: %d\n", f(2))
	fmt.Print("\n")
}
