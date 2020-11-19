package demo

import "fmt"

func Datatype() {
	/*
	 * 数字类型
	 * uintptr 	: uint32(32bit-sys) uint64(64bit-sys)
	 * int 		: int32(32bit-sys) int64(64bit-sys)
	 * uint 	: uint32(32bit-sys) uint64(64bit-sys)
	 *
	 * int8 uint8(byte)
	 * int16 uint16
	 * int32(rune) uint32
	 * int64 uint64
	 * float32 float64 		: 浮点数
	 * complex64 complex128 : 虚数
	 */
	var c byte = 'A'
	var n int = 0xAA
	var f float32 = 3.14
	fmt.Println("c=", c)
	fmt.Println("n=", n)
	fmt.Println("f=", f)

	var b bool
	b = true
	fmt.Println("b=", b)

	var s string = "C:\\"
	fmt.Println("s=", s)
	var ss = `C:\`
	fmt.Println("ss=", ss)

	fmt.Print("\n")
}

func ComplexDatatype() {
	// 数组
	var arr1 [10]byte
	var arr2 = [10]byte{7, 6, 5, 4, 3, 2, 1, 'A', 'B', 'C'}
	arr1 = arr2
	fmt.Println("arr1=", arr1)

	// 指针
	var b byte
	var pByte *byte = &b
	*pByte = 233
	fmt.Println("TestPointer b =", b)

	// 数组指针
	var arr3 = [2]byte{11, 22}
	var arr4 = arr3      // 深拷贝数组
	var pByteArr = &arr4 // 浅拷贝数组
	pByteArr[0] = 111
	pByteArr[1] = 222
	fmt.Println("TestPointer arr3=", arr3)
	fmt.Println("TestPointer arr4=", arr4)

	// 空指针
	var pn *int
	fmt.Println("TestPointer pn=", pn == nil)

	// 结构体
	type Book struct {
		id     int
		title  string
		author string
	}
	var book1 = Book{123, "Thinking in C ", "?"}
	book2 := Book{id: 321, title: "Thinking in Go", author: "!"}

	var pBook1 = &book1
	pBook2 := &book2

	fmt.Println("struct book1 =", book1, pBook1.id)
	fmt.Println("struct book2 =", book2, pBook2.id)
	fmt.Print("\n")
}
