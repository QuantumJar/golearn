package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}
type API int

//一个函数可以作为RPC方法的要求：
//1.他必须是一个有接受者的方法，并且是一个导出的方法
//2.他必须有两个参数，这两个参数都是导出类型,或者是内置类型
//3.返回值类型是error

var database []Item

func main() {
	// fmt.Println("初始数据库:", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("第二个数据库:", database)

	// DeleteItem(b)
	// fmt.Println("第3个数据库:", database)
	//
	var api = new(API)

	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering api", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("listener error", err)
	}

	log.Printf("serving rpc on port 4040")
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving", err)
	}

}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Title == title {
			getItem = val
			break
		}
	}
	//
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for i, val := range database {
		if val.Title == edit.Title {
			database[i] = edit
			changed = edit
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for i, val := range database {
		if val.Title == item.Title {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	//reply 是一个指针，*reply 表示获取这个指针指向的值
	*reply = del
	return nil
}

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}
