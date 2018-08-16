package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱拜仁!"
	for _, b := range []byte(s) {
		fmt.Printf("%X", b)

	}
	fmt.Println(len(s))

	for i, ch := range s {
		fmt.Printf("%d, %X", i, ch)
	}
	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, char := range []rune(s) {
		fmt.Printf("%d, %c", i, char)
	}

}
