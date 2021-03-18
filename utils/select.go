package utils

import (
	"fmt"
	"time"
)

func Select() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "msg1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "msg2"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)

	case <-time.After(time.Second):
		fmt.Println("timeout")
		//default:
		//	fmt.Println("Default")
	}
}
