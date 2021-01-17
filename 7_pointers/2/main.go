package main

import "fmt"

type Person struct {
	name string
}

func changeMe(pPerson *Person) {
	pPerson.name = fmt.Sprintf("changed_%s", pPerson.name)
}

func main() {
	person := Person{"Harry"}
	fmt.Println(person)
	changeMe(&person)
	fmt.Println(person)

}
