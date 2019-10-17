package piscine

import (
	//"github.com/01-edu/z01"
	"fmt"
)

func PrintStr(str string) {
	for _, gh := range str {
		fmt.Println("%a", gh)
	}
}
