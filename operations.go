package main

import (
	"fmt"
	"strings"
)

func createDevice() string {
	fmt.Println("***** Create device *****")

	fmt.Println("Ange namn:")
	var newName string
	fmt.Scanln(&newName)

	return newName
}

func listAll() {
	fmt.Println("***** Listing all *****")
	for _, varde := range devices {
		fmt.Printf("%s\n", varde)
	}
}

func updateDevice() {
	fmt.Println("***** Update device *****")
	for index, varde := range devices {
		fmt.Printf("%d:%s\n", index+1, varde)
	}
	sel := inputIntData(1, len(devices), "Ange vilket nummer du vill uppdatera:")

	var nyttVarde string
	fmt.Println("Ange nytt namn:")
	fmt.Scanln(&nyttVarde)
	devices[sel-1] = nyttVarde
}

func searchDevice() {
	fmt.Println("***** Search device *****")
	fmt.Printf("Ange sök:")
	var ord string
	fmt.Scanln(&ord)

	for _, varde := range devices {
		searchLower := strings.ToLower(ord)
		vardeLower := strings.ToLower(varde)

		if strings.Contains(vardeLower, searchLower) {
			fmt.Printf("%s\n", varde)
		}
	}

}
