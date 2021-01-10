package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float32
}

func (rcv Square) Area() float32 {
	return rcv.length * rcv.length
}

type Circle struct {
	radius float32
}

func (rcv Circle) Area() float32 {
	return math.Pi * rcv.radius * rcv.radius
}

type Shape interface {
	Area() float32
}

func info(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	circle := Circle{
		radius: 1.5,
	}
	square := Square{
		length: 3.0,
	}
	info(circle)
	info(square)

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
