package demo

import (
	"fmt"
	"time"
)

func Thread() {
	var say = func(name string) {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s-%d ", name, i)
		}
	}
	fmt.Print("GoRoutine: ")
	go say("A")
	go say("B")
	say("O")
}
