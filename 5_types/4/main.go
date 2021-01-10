package main

import "fmt"

func main() {
	anon := struct {
		blub map[string]string
		blab []int
	}{
		blub: map[string]string{
			"first":  "firstValue",
			"second": "secondValue",
		},
		blab: []int{
			1,
			2,
		},
	}

	for k, v := range anon.blub {
		fmt.Println(k, v)
	}

}
