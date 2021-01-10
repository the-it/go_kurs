package main

import "fmt"

func main() {
	intArray := []int{1, 2, 3, 4}
	fmt.Println(foo(intArray...))
	fmt.Println(bar(intArray))

}

func foo(variadic ...int) int {
	sum := 0
	for _, i := range variadic {
		sum += i
	}
	return sum
}

func bar(intArray []int) int {
	sum := 0
	for _, i := range intArray {
		sum += i
	}
	return sum
}
