package utils

import (
	"fmt"
	"math/rand"
)

func Switch()  {
	var i int = 0;
	i = rand.Int()%2;
	switch i {
	case 0:
		fmt.Println("偶数")
		break;
	case 1:
		fmt.Println("奇数");
		break;
	default:
		fmt.Println("语法错误");
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
