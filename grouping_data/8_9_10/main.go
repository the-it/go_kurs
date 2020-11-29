package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	x["myself"] = []string{`not, james bond`, "something else"}
	printMap(x)
	delete(x, "no_dr")
	printMap(x)

}

func printMap(x map[string][]string) {
	for key, value := range x {
		fmt.Printf("%v favourite things:\n", key)
		for idx, item := range value {
			fmt.Printf("\t %v: %v\n", idx, item)
		}
	}
	fmt.Println()
}
