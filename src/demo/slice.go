package demo

import "fmt"

func Slice() {
	// allocate
	var arr = make([][]int, 5, 5)
	fmt.Println("arr =", arr)
	arr = [][]int{
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

	fmt.Printf("len=%d cap=%d\n", len(all), cap(all))
	fmt.Println("all  =", all)
	fmt.Println("row1 =", row1)
	fmt.Println("row2 =", row2)
	fmt.Println("row3 =", row3)
	fmt.Print("\n")
}

func Range() {
	var arr = [][]int{
		{11, 12, 13, 14, 15},
		{21, 22, 23, 24, 25},
		{31, 32, 33, 34, 35},
		{41, 42, 43, 44, 45},
		{51, 52, 53, 54, 55},
	}
	fmt.Println("arr  =", arr)
	for _, row := range arr {
		fmt.Println("row =", row)
	}
	for num, row := range arr {
		fmt.Println(num, "=", row)
	}
	fmt.Print("\n")
}
