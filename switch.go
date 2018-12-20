package main

import (
	"fmt"
	"time"
)

func main(){
	i:= 2
	fmt.Print("Write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")

	}
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		default:
			fmt.Println("not known%T\n",t)

		}
	}

	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("hey")
}

