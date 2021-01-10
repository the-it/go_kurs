package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	sedanObject := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "yellow",
		},
		luxury: false,
	}

	fmt.Println(sedanObject.doors)
	fmt.Println(sedanObject)

}
