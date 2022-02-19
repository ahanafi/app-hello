package main

import (
	"fmt"
	hello_go "github.com/ahanafi/hello-go"
)

func main() {
	hello := hello_go.SayHello()
	fmt.Println(hello)
}
