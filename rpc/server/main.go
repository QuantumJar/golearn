package  main

import (
	"fmt"
	"net/rpc"
)

func main(){

	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	defer  client.Close()
	 args := []int{2,3}
	var reply int
	client.Call("mathService.Multiply",args,&reply)

	fmt.Println(reply)

}