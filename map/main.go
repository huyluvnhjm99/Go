package main

import "fmt"

func main() {

	//var pokemi map[string]string
	// fmt.Println(pokemu)
	// delete(pokemu, 5)
	// fmt.Println(pokemu)
	//pokemu := make(map[int]string)

	pokemonMap := map[string]string{
		"pikachuu": "001",
		"pikachii": "002",
	}

	printMap(pokemonMap)
}

func printMap(m map[string]string) {
	for pokemon, code := range m {
		fmt.Println("The code of '" + pokemon + "' is: " + code)
	}
}
