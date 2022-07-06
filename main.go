package main

import (
	"fmt"
)

func printMaterials(materials ItemList) {
	for material, amount := range materials {
		fmt.Printf("%s: %f\n", material, amount)
	}
	fmt.Println("------------------------------")
}

func main() {

	for {
		itemName := ""
		itemAmount := 0.0
		fmt.Scanf("%s %f", &itemName, &itemAmount)
		if itemName == "" || itemAmount == 0 {
			break
		}
		materials := Materials(itemName, itemAmount)
		printMaterials(materials)
		_, done := Simplify(materials)
		for !done {
			materials, done = Simplify(materials)
			printMaterials(materials)
		}
		fmt.Println()
	}
}
