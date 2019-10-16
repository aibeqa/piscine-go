package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for o := '0'; o <= '9'; o++ {
		for v := '0'; v <= '9'; v++ {
			for f := '0'; f <= '9'; f++ {
				for e := '0'; e <= '9'; e++ {
					if o == '9' && v == '8' && f == '9' && e == '9' {
						z01.PrintRune(o)
						z01.PrintRune(v)
						z01.PrintRune(32)
						z01.PrintRune(f)
						z01.PrintRune(e)
						z01.PrintRune(10)
						break
					}
					if o == f && v < e || o < f {
						z01.PrintRune(o)
						z01.PrintRune(v)
						z01.PrintRune(32)
						z01.PrintRune(f)
						z01.PrintRune(e)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}
