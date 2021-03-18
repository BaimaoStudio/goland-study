package utils

import (
	"fmt"
	"time"
)

func Channel() {
	msg := make(chan string)
	go func() {
		msg <- "ping"
		fmt.Println("pong")
	}()
	go func() {
		msg2 := <-msg
		fmt.Println(msg2)
	}()
	time.Sleep(time.Second)

	//channel buffer
	msg3 := make(chan string, 2)
	msg3 <- "channel"
	msg3 <- "buffer"
	fmt.Println(<-msg3)
	fmt.Println(<-msg3)
	done := make(chan bool)
	go sychronization(done)
	fmt.Println(<-done)
}

func sychronization(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
