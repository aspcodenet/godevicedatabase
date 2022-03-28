package main

import "fmt"

func main() {

	for {
		showMenu()
		var selection int

		selection = inputIntData(1, 5, "Val")

		switch selection {
		case 1:
			fmt.Println("Create device")
		case 2:
			fmt.Println("List all")
		case 3:
			fmt.Println("Update device")
		case 4:
			fmt.Println("Sök device")
		case 5:
			//break would just break the switch!
			return

		}
	}
	// Kom ihåg lite rast också!
	//SES kl 14:00 !
	// Då kör vi  funktioner
	// och refaktoriserar dessutom denna kod
}
