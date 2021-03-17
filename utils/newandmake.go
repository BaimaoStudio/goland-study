package utils

import "fmt"

/**
 *	关于new和make的基础原理和用法
 *  make 被用来分配引用类型的内存： map, slice, channel
 *	new 被用来分配除了引用类型的所有其他类型的内存： int, string, array等
 */
func Newandmake()  {

	var i *int = new(int)
	fmt.Println(*i)

	s := make([]string, 3);
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	c := make([]string,3);
	copy(c,s)
	fmt.Println(c,s)
}
