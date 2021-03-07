package main

import (
	"fmt"
	"time"
)

const AUTHOR = "Tanay Mishra"

func main() {
	fmt.Println("Hello" + "Tanay, You are")

	var a = "Variable defined at initialisation doesn`t need any type"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println("You can define more than one var in single statements like b = ", b, " and c = ", c)
	var d int
	fmt.Println(d)
	e := "Short declaration used"
	fmt.Println(e)

	fmt.Println("Author is ", AUTHOR)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even number")
		} else {
			fmt.Println(i, "is odd number")
		}
	}

	if abc := 2; abc%2 == 0 {
		fmt.Println(abc, "is even")
	} else {
		fmt.Println(abc, "is odd")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is weekday")
	}
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Bool type")
		case int:
			fmt.Println("INT Type")
		default:
			fmt.Println("default")
		}
	}

	whatAmI(true)
	whatAmI(2)

}
