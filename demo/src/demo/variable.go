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
