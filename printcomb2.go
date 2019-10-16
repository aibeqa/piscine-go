package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var count int = 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i < j2 {
				z01.PrintRune(44)
				z01.PrintRune(32)
				z01.PrintRune("%02d %02d", i, j)
				count++
			}
		}
	}
	z01.PrintRune(10)
}
