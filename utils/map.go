package utils

import "fmt"

func Map()  {
	//map可以不定义容量
	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m["key1"])
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["key3"])
	m["key3"] = 3
	fmt.Println(m["key3"])
	//删除
	delete(m,"key2");
	//prs是否存在值 a当前值
	a,prs := m["key2"]
	fmt.Println(a)
	fmt.Println(prs)
	//初始化map
	n := map[string]int{"key1":1,"key2":2}
	fmt.Println(n)
}
