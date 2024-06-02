package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个string类型的channel
	strChannel := make(chan string)

	//声明一个匿名函数并且在另一个线程中运行他
	go func() {
		//数据写入channel使用了箭头函数
		strChannel <- "write something to the string channel"
		fmt.Println("我已经发送了数据，我可以退出了吗")
	}()

	//在主线程中读取strChannel中的数据，
	time.Sleep(time.Second * 5)
	//join point of this fork join model
	data := <-strChannel

	fmt.Println(data)
}
