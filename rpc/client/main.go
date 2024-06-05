package main

import (
	"fmt"
	"net/rpc"
)

type ExampleArg struct {
	X int
}

type ExampleReply struct {
	Y int
}

func main() {

	args := ExampleArg{}
	args.X = 10
	reply := ExampleReply{}

	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	client.Call("MathService.Add", args, &reply)

	fmt.Printf("结果: %v", reply)
}
