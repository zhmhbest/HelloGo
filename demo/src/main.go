package main

// 导入一个包: import (子包 目录)
// 导入所有包: import (目录)
import hello "./demo"

func main() {
	hello.SayHello()
}