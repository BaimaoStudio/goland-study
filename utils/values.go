package utils

import "fmt"
func Value()  {
	var abc int = 100
	var bcd int = 200
	cde := abc+bcd
	var temp string = "100+200="
	fmt.Println(temp,cde)

	var boo1 bool = true
	var boo2 bool = false

	fmt.Println(boo1 || boo2)
	fmt.Println(boo1 && boo2)
	fmt.Println(!boo1)

}
