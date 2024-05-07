package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	i := r.Intn(100)
	for {
		fmt.Print(i)
	}
}
