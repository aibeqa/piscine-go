package piscine

import (
	//"github.com/01-edu/z01"
	"fmt"
)

func PrintStr(str string) {
	for _, letter := range str {
		fmt.Println("%a", letter)
	}
}
