package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Person struct {
	Name string
	Age  int
}

type emp []Person

type Request struct {
	Data interface{}
}

type Response struct {
	Data interface{}
	Code int
}

func main() {
	var response Response

	// var employeeList emp

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection err:", err)
	}
	arg := Request{}
	client.Call("Person.GetAllEmployees", arg, &response)

	fmt.Printf("response %v", response.Code)
}
