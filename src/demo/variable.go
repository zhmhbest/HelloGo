package demo

import "fmt"

func Define() {
	var i1 int = 1           // 声明变量类型并赋值
	var b1, b2 byte = 1, 2   // 两个相同类型的变量
	b1, b2 = 0xf, 0xf        // 已声明的变量进行赋值
	var f1, f2 = 3.14, 0.618 // 自动判断变量类型
	s1, s2 := "str1", "str2" // 自动创建并赋值变量（仅限函数内使用）
	fmt.Println(i1, b1, b2, f1, f2, s1, s2)
	fmt.Print("\n")
	fmt.Print("\n")
}

func Constant() {
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
	fmt.Print("\n")
}
