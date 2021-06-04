package main

import "fmt"

type Person struct {
	name string
}

func (rcv *Person) speak(){
	fmt.Println(rcv.name)
}

type Human interface {
	speak()
}

func saysomething (h Human){
	h.speak()
}

func main() {
	person := Person{"Harry"}
	saysomething(&person)

}
