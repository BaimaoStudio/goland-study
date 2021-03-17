package utils

import "fmt"

func Func1()  {
	a,b := func2(1,2)
	fmt.Println(a,b)
	_,c := func2(1,2)
	fmt.Println(c)
	func3(1,2,3,4,5,6)

	f4 := func4()
	fmt.Println(f4())
	fmt.Println(func4())

	fmt.Println(func5(3))
}

func func2(a int,b int)(int,int){
	return a+b,a-b
}

func func3(num ... int)  {
	fmt.Println(num)
	for _,v := range num{
		fmt.Println(v)
	}
}

//闭包
func func4() func() int {
	i:=0
	return func() int {
		i++
		return i
	}
}

//递归
func func5(i int) int {
	if i>5 {
		return 1
	}
	return func5(i+1)
}
