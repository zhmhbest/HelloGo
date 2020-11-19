package demo

import (
	"errors"
	"fmt"
)

func Exception() {
	var f = func() (float32, error) {
		// 自定义异常
		return 3.1415, errors.New("自定义异常")
	}

	// 捕获异常
	var pi, err = f()
	if err != nil {
		fmt.Println("Exception: ", err.Error())
	} else {
		fmt.Println("Exception: NoError", pi)
	}
	fmt.Print("\n")
}
