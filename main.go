package main

import (
	"fmt"
	hello_go "github.com/ahanafi/hello-go/v2"
)

func main() {
	hello := hello_go.SayHello("Hanafi!")
	fmt.Println(hello)
}
