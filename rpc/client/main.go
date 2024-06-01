package  main

import (
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct {

}

func (m *MathService)Multiply(args [2]int,reply *int)  error{

	*reply = args[0]*args[1]
	return  nil
}
func main(){

	var mathService = new(MathService)

	rpc.Register(mathService)

	rpc.HandleHTTP()

	listener, _ := net.Listen("tcp", ":1234")


	http.Serve(listener,nil)

}