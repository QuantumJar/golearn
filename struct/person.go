package main

import "fmt"

type Person struct {
	firstname   string
	lastname    string
	contactInfo ContactInfo
}

type ContactInfo struct {
	email   string
	address string
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func (personPtr *Person) setFirstname(firstname string) {
	// 这种写法和下面的写法有什么不一样呢？
	(*personPtr).firstname = firstname

	// personPtr.firstname = firstname
}
