package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

type ExampleArg struct {
	X int
}

type ExampleReply struct {
	Y int
}

func (m *MathService) Add(args ExampleArg, reply *ExampleReply) error {
	reply.Y = args.X + args.X
	return nil
}
func main() {
	var m = new(MathService)
	rpc.Register(m)

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, nil)

}
