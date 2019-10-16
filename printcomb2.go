package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var count int = 0

	for num1 := 0; num1 < 100; num1++ {
		for num2 := 0; num2 < 100; num2++ {
			if num1 < num2 {
				z01.PrintRune(44)
				z01.PrintRune(32)
				fz01.PrintRune("%02d %02d", num1, num2)
				count++
			}
		}
	}
	z01.PrintRune()
}
