package main

import "fmt"

func main() {
	go someFunc("somefunc")
	fmt.Println("print from main routine")
}

func someFunc(str string) {
	fmt.Println(str)
}
