package utils

import "fmt"

/**
	slice 切片 array数组

 */
func Slice()  {
	slice := make([]string,5);
	slice[0] = "1"
	slice[1] = "2";
	slice[2] = "3";
	l := slice[2:5]
	fmt.Println(l)

	l3 := slice[:4];

	fmt.Println(l3)

	l4 := slice[2:]
	fmt.Println(l4);

	//slice[9] = "4";
	fmt.Println(slice)
	l2 := append(slice, "append")
	fmt.Println(l2)

	l5 := append(slice,"ap1","ap2")

	fmt.Println(l5)
	//使用make创建切片不会置零
	arr := make([][]int,5,5)
	for _,value := range arr {
		fmt.Println(value)
	}
}
