package main

import "fmt"

//声明一个接口a和它的函数们
type a interface {
	methodA() string
}

type b struct {
}

//声明了一个receiver是b类型的methodA,这个方法和接口a中的方法签名一致，这里就隐式的表示了b类型也是a类型（a是一个接口）
func (b) methodA() string {
	return "b类型自己的实现"
}

type c struct{}

func (c) methodA() string {
	return "c类型自己的实现"
}

//这个方法的形参是接口a，这个接口a并没有一个具体的methodA的实现
func someMethod(a a) {
	fmt.Println(a.methodA())
}

func main() {
	b := b{}
	c := c{}

	someMethod(b)
	someMethod(c)
}
