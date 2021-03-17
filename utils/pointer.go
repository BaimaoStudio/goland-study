package utils

import "fmt"

func Pointer()  {

	var a *int = new(int)

	fmt.Println(a)
	fmt.Println(*a)
	b := 1
	a = &b
	fmt.Println(a)
	fmt.Println(*a)

	//var c *int
	//*c = 10
	//fmt.Println(*c)
	//fmt.Println(c)
}
