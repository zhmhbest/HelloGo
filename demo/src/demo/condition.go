package demo

import "fmt"

func Condition() {
	// if
	if true {
		fmt.Println("if true")
	} else if false {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	// switch
	var x = 0
	switch x {
	case 0:
		fmt.Println("switch", 3)
		break
	default:
		fmt.Println("switch default")
	}

	// for
	for i := 0; i < 3; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Print("\n")

	// death-loop
	for {
		fmt.Println("death-loop")
		break
		//continue
	}
	fmt.Print("\n")
}
