package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteIceCream []string
}

func main() {
	personOne := person{
		firstName: "A",
		lastName:  "Aa",
		favouriteIceCream: []string{
			"erste Sorte",
			"zweite Sorte",
		},
	}
	for _, iceCream := range personOne.favouriteIceCream {
		fmt.Println(iceCream)
	}
}
