package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "2"},
		{"Jes", "Bd", "3"},
	}
	for _, itemList := range x {
		for _, item := range itemList {
			fmt.Println(item)
		}
	}

}
