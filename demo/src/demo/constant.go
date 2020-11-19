package demo

import "fmt"

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
