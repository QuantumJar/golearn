package main

import (
	"fmt"
	"sort"
)

/**

**/

type interfaceA interface {
	MethodA() string
	A()
}

type interfaceB interface {
	A()
	MethodB() string
}

type Person struct {
	name string
}

func (p Person) A() {
	fmt.Println("person")
}

func (p Person) MethodA() string {
	return "MethodA"
}
func (p Person) MethodB() string {
	return "MethodB"
}

func main() {
	var intList []int = []int{1, 2, -1, 92, 31}

	sort.Ints(intList)

	fmt.Print(intList)

}
