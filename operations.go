package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createDevice() *Device {
	fmt.Println("***** Create device *****")

	fmt.Println("Ange namn:")
	var newName string
	fmt.Scanln(&newName)

	fmt.Println("Ange vikt i gram:")
	var weight int
	fmt.Scanln(&weight)

	fmt.Println("Ange Manufacturer:")
	var manufacturer string
	fmt.Scanln(&manufacturer)

	return &Device{Name: newName, WeightInGrams: weight, Manufacturer: manufacturer}
}

func listAll(devices []Device) {
	fmt.Println("***** Listing all *****")
	for _, varde := range devices {
		fmt.Println(varde)
	}
}

func updateDevice(devices []Device) {
	fmt.Println("***** Update device *****")
	for index, varde := range devices {
		fmt.Printf("%d:%s\n", index+1, varde.Name)
	}
	sel := inputIntData(1, len(devices), "Ange vilket nummer du vill uppdatera:")

	var nyttVarde string
	fmt.Printf("Ange nytt namn (eller blank för att ha kvar '%s'):", devices[sel-1].Name)
	fmt.Scanln(&nyttVarde)
	if len(nyttVarde) > 0 {
		devices[sel-1].Name = nyttVarde
	}

	fmt.Printf("Ange ny manufacturer (eller blank för att ha kvar '%s'):", devices[sel-1].Manufacturer)
	fmt.Scanln(&nyttVarde)
	if len(nyttVarde) > 0 {
		devices[sel-1].Manufacturer = nyttVarde
	}

	fmt.Printf("Ange ny vikt (eller blank för att ha kvar '%d'):", devices[sel-1].WeightInGrams)
	fmt.Scanln(&nyttVarde)
	if len(nyttVarde) > 0 {
		intVar, _ := strconv.Atoi(nyttVarde)
		devices[sel-1].WeightInGrams = intVar
	}
}

func searchDevice(devices []Device) {
	fmt.Println("***** Search device *****")
	fmt.Printf("Ange sök:")
	var ord string
	fmt.Scanln(&ord)

	for _, device := range devices {
		searchLower := strings.ToLower(ord)
		nameLower := strings.ToLower(device.Name)
		manufacturerLower := strings.ToLower(device.Manufacturer)

		if strings.Contains(nameLower, searchLower) || strings.Contains(manufacturerLower, searchLower) {
			fmt.Println(device)
		}
	}

}
