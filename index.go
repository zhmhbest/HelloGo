package main

import (
	"fmt"
	"time"
)
import "errors"

// 单行注释
/* 多行注释 */

func hello() {
	fmt.Print("Hello ")
	println("World!")
}

func helloVariable() {
	var i1 int = 10
	var b1, b2 byte = 1, 2   // 两个相同类型的变量
	b1, b2 = 0xa, 0xf        // 已声明的变量进行赋值
	var f1, f2 = 3.14, 0.618 //自动判断变量类型
	s1, s2 := "str1", "str2" //自动判断变量类型（仅限函数内使用） 创建并赋值变量
	fmt.Println(i1, b1, b2, f1, f2, s1, s2)
}

func helloConstant() {
	const GF = 0.618
	const PI float64 = 3.1415926535897
	fmt.Println("GF=", GF)
	fmt.Println("PI=", PI)
	// 独立计数器常量 iota : const 之后每此自增1
	const (
		//i0 //此处报错
		i1 = "*"
		i2 = iota
		i3
		i4
		i5 = "#"
		i6
		i7
		i8 = iota
		i9
	)
	fmt.Println("iota1:", i1, i2, i3, i4, i5, i6, i7, i8, i9)
	// 利用iota位移实现位运算
	const (
		bin1 = 1 << iota
		bin2 = 1 << iota
		bin3 = 1 << iota
		bin4 = 1 << iota
		bin5 = 1 << iota
		bin6 = 1 << iota
		bin7 = 1 << iota
		bin8 = 1 << iota
	)
	fmt.Println("iota2:", bin1, bin2, bin3, bin4, bin5, bin6, bin7, bin8)
}

func helloSimpleDatatype() {
	/*
	 * 数字类型
	 * uintptr : uint32(32bit-sys) uint64(64bit-sys)
	 * int : int32(32bit-sys) int64(64bit-sys)
	 * uint : uint32(32bit-sys) uint64(64bit-sys)
	 * int8 uint8(byte)
	 * int16 uint16
	 * int32(rune) uint32
	 * int64 uint64
	 * float32 float64 : 浮点数
	 * complex64 complex128 : 虚数
	 */
	var c byte = 'A'
	var n int = 0xAA
	var f float32 = 3.14
	fmt.Println("c=", c)
	fmt.Println("n=", n)
	fmt.Println("f=", f)

	/*
	 * 布尔类型
	 * bool = true | false
	 */
	var b bool
	b = true
	fmt.Println("b=", b)

	/*
	 * 布尔类型
	 * string = true | false
	 */
	var s string = "字符串！"
	fmt.Println("s=", s)
}

func helloComplexDatatype() {
	/*
	 * 数组类型
	 */
	var bs1 [10]byte
	var bs2 = [10]byte{7, 6, 5, 4, 3, 2, 1, 'A', 'B', 'C'}
	bs1 = bs2
	fmt.Println("byte[10]=", bs1)

	/*
	 * 指针类型
	 */
	var b byte
	var bp *byte = &b
	*bp = 233
	fmt.Println("Pointer b =", b)
	// 数组
	var bs [2]byte
	var p = &bs
	p[0] = 111
	p[1] = 222
	fmt.Println("Pointer b[2] =", bs)
	// 空指针
	var pn *int
	fmt.Println("Pointer pn =", pn == nil)

	/*
	 * 结构体
	 */
	type Book struct {
		id     int
		title  string
		author string
	}
	var book1 Book = Book{123, "Title", "Author"}
	book2 := Book{id: 321, title: "T", author: "A"}
	var pbook1 *Book = &book1
	pbook2 := &book2
	fmt.Println("struct book1 =", book1, pbook1.id)
	fmt.Println("struct book2 =", book2, pbook2.id)
}

func helloControl() {
	// if
	if true {
		fmt.Println("if true")
	} else if false {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
	// for
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	// death-loop
	for {
		fmt.Println("death-loop")
		break
		//continue
	}
}

type funVar func() int //定义函数类型
func helloCloser() funVar {
	var count = 0 //闭包变量
	f := func() int {
		count++
		return count
	}
	return f
}

func helloSlice() {
	var slice = make([][]int, 5, 5) //allocates
	fmt.Println("slice =", slice)

	var arr = [][]int{
		{11, 12, 13, 14, 15},
		{21, 22, 23, 24, 25},
		{31, 32, 33, 34, 35},
		{41, 42, 43, 44, 45},
		{51, 52, 53, 54, 55},
	}
	all := arr[:]    // 全切
	row1 := arr[:3]  // 行切 *~3行
	row2 := arr[1:4] // 行切 2~4行
	row3 := arr[3:]  // 行切 4~*行

	fmt.Printf("len=%d cap=%d slice=%v\n", len(all), cap(all), all)
	fmt.Println("row1 =", row1)
	fmt.Println("row2 =", row2)
	fmt.Println("row3 =", row3)
}

func helloRange() {
	var arr = [][]int{
		{11, 12, 13, 14, 15},
		{21, 22, 23, 24, 25},
		{31, 32, 33, 34, 35},
		{41, 42, 43, 44, 45},
		{51, 52, 53, 54, 55},
	}
	fmt.Println(arr)
	for _, num1 := range arr {
		fmt.Println("v =", num1)
	}
	for i, v := range arr {
		fmt.Println(i, "=", v)
	}
}

func helloMap() {
	var kv map[string]string = make(map[string]string)
	kv["key1"] = "value1"
	kv["key2"] = "value2"
	kv["key3"] = "value3"
	delete(kv, "key2")
	for i, v := range kv {
		fmt.Println(i, "=", v)
	}
}

func helloException() {
	f := func() (float32, error) {
		// 自定义异常
		return 0.618, errors.New("自定义异常")
	}
	// 捕获异常
	var GR, err = f()
	if err != nil {
		fmt.Println("Exception", err.Error())
	}
	fmt.Println("Exception", GR)
}

func helloThread() {
	say := func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(s)
		}
	}
	fmt.Print("GoRoutine: ")
	go say("-")
	say(".")
}

func main() {
	//helloPackage()
	hello()
	helloVariable()
	helloConstant()
	helloSimpleDatatype()
	helloComplexDatatype()
	helloControl()
	var f = helloCloser()
	fmt.Println(f())
	fmt.Println(f())
	helloSlice()
	helloRange()
	helloMap()
	helloException()
	helloThread()
}
