package utils

import (
	"fmt"
	"time"
)

func Goroutines() {
	test1()
}

func func1(routines string) {
	for i := 1; i <= 3; i++ {
		fmt.Println("routine:", routines, i)
	}
}

func test1() {
	go func1("Test1")
	go func1("Test2")

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("routine:Test3", i)

		}
	}()
	//不sleep会导致 go未开始执行 main函数结束 所有goroutines中断
	time.Sleep(time.Second)
}
