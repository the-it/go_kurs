package main

import "fmt"

func main() {
	x := []int{0, 0, 0, 0, 0}
	for i := 1; i <= 5; i++ {
		x[i-1] = i
	}
	fmt.Println(x[1:5])
}
