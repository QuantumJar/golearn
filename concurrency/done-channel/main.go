package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go boring(done)

	//主线程3秒后给done channel 发信号
	time.Sleep(time.Second * 3)

	done <- false

	time.Sleep(time.Hour * 3)

}

func boring(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("收到done信号")
			return
		default:
			fmt.Println("working")
		}
	}
}
