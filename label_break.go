package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")

Exit:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i+j > 15 {
				fmt.Println("exit")
				break Exit
			}
		}
	}

	fmt.Println("3")
}
