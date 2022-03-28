package main

import "fmt"

func inputIntData(min int, max int, prompt string) int {
	for {
		var val int
		fmt.Print("Val:")
		_, err := fmt.Scanln(&val)
		if err != nil {
			var clearbuf string
			fmt.Scanln(&clearbuf)
			fmt.Println("Inte valid nummer")
			continue
		}
		if val < min || val > max {
			fmt.Println("Mata in ett tal mellan 1 och 5 tack")
			continue
		} else {
			return val
		}
	}
}
