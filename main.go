package main

import "fmt"

func main() {

	for {
		var selection int

		fmt.Println("1. Create device")
		fmt.Println("2. List all")
		fmt.Println("3. Update device")
		fmt.Println("4. Sök")
		fmt.Println("5. Exit")

		fmt.Print("Val:")
		_, err := fmt.Scanln(&selection)
		if err != nil {
			var clearbuf string
			fmt.Scanln(&clearbuf)
			fmt.Println("Inte valid nummer")
			continue
		}

		if selection < 1 || selection > 5 {
			fmt.Println("Ange ett tal mellan ett och fem tack")
			continue
		}

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

}
