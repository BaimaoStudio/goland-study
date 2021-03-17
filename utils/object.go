package utils

import "fmt"

type Object struct {
	name string
	age int
}

func Obj()  {
	p := Object{"啊啊",1}
	fmt.Println(p)
	fmt.Println(p.age,p.name)
	fmt.Println(&p)
	sp := p
	p.age = 2
	fmt.Println(sp)
	sp2 := &p
	p.age = 3
	fmt.Println(sp2)
}
