package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.bilibili.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	bs := make([]byte, 100)
	resp.Body.Read(bs)
	fmt.Print(string(bs))
}
