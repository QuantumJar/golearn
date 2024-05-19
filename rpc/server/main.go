package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Person struct {
	Name string
	Age  int
}

var employeeList []Person

type Request struct {
	Data interface{}
}

type Respose struct {
	Data interface{}
	Code int
}

func main() {
	p1 := Person{Name: "zhang", Age: 30}
	p2 := Person{Name: "tom", Age: 25}
	p3 := Person{Name: "jack", Age: 20}

	employeeList = append(employeeList, p1)
	employeeList = append(employeeList, p2)
	employeeList = append(employeeList, p3)

	fmt.Printf("员工：%v", employeeList)
	// arith := new(Arith)
	// rpc.Register(arith)
	// rpc.HandleHTTP()
	// l, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("listen error:", err)
	// }
	// go http.Serve(l, nil)
	person := new(Person)
	rpc.Register(person)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(l, nil)

	fmt.Println("服务开启...")

}

//按照go RPC的约定，一个方法可以作为RPC方法要满足一些条件
// 1.the method's type is exported.
// 2.the method is exported.
// 3.the method has two arguments, both exported (or builtin) types.
// 4.the method's second argument is a pointer.
// 5.the method has return type error.

// 这里使用request和response两个参数来封装请求参数和返回值。
func (p *Person) AddEmployee(request Request, response *Respose) error {

	//这里Data是一个Person，但是如何类型转换呢？
	//使用类型断言来判断
	personToAdd, ok := request.Data.(Person)
	if !ok {
		return errors.New("非法参数！参数不是Person类型！！")
	}

	employeeList = append(employeeList, personToAdd)

	response.Data = personToAdd
	response.Code = 200
	return nil
}

// func (p *Person) GetAllEmployees(request Request, response *Respose) error {

// }
