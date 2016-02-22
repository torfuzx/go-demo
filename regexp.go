package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validPhone = regexp.MustCompile(`^\+?[0-9][0-9]*$`)
	fmt.Println(validPhone.MatchString("0454353451232132132132132123"))
}
