package main

import "fmt"

type englishBot struct{}
type chineseBot struct{}

// 声明了一个bot的借口，是不是任何同时拥有这个方法实现的结构都属于bot?
type bot interface {
	getGreeting() string
}

// getGreeting肯定是绑定在各自的类型上的
func (englishBot) getGreeting() string {
	return "hello there"
}

func (chineseBot) getGreeting() string {
	return "你好"
}

//声明了一个参数是bot的方法，似乎这个b和
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	eb := englishBot{}
	cb := chineseBot{}

	printGreeting(eb)
	printGreeting(cb)
}
