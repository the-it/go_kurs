package main

import (
	"fmt"
)

func printer(a int) int {
	fmt.Println("Hello")
	return 5
}

func caller(functionToCall func(int) int) {
	functionToCall(4)
}

func main() {
	caller(printer)

}
