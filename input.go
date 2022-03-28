package main

import "fmt"

func inputIntData(min int, max int, prompt string) int {
	for {
		var val int
		fmt.Print(prompt)
		_, err := fmt.Scanln(&val)
		if err != nil {
			var clearbuf string
			fmt.Scanln(&clearbuf)
			fmt.Println("Inte valid nummer")
			continue
		}
		if val < min || val > max {
			fmt.Printf("Mata in ett tal mellan %d och %d tack\n", min, max)
			continue
		} else {
			return val
		}
	}
}
