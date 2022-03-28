package main

import "fmt"

func main() {
	var selection string

	fmt.Println("1. Create device")
	fmt.Println("2. List all")
	fmt.Println("3. Update device")
	fmt.Println("4. Sök")
	fmt.Println("5. Exit")

	fmt.Print("Val:")
	fmt.Scanln(&selection)

	fmt.Printf("Du valde:%s\n", selection)
}
