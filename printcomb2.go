package piscine

import "fmt"

func PrintComb2() {
	var count int = 0

	for num1 := 0; num1 < 100; num1++ {
		for num2 := 0; num2 < 100; num2++ {
			if num1 < num2 {
				if count > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%02d %02d", num1, num2)
				count++
			}
		}
	}
	fmt.Println()
}