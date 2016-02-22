package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := `asd123 我是谁`

	fmt.Println("len--", len(s))
	fmt.Println("runelen--", len([]rune(s)))
	fmt.Println("RuneCountInString--", utf8.RuneCountInString(s))
}
