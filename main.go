package main

import (
	"fmt"
	"go_learn/user"
)

func main() {
	s := user.Hello()

	fmt.Printf("s=%s", s)
}
