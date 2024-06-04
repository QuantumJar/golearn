package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan string)

	go boring("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("you say : %q\n", <-c)
	}
	fmt.Println("you are boring , i'm leaving")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
