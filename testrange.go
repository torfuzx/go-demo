package main

import "fmt"

func generator(strings chan string) {
	strings <- "Five hour's New York jet lag"
	strings <- "and Cayce Pollard wakes in Camden Town"
	strings <- "to the dire and ever-decreasing circles"
	strings <- "of disrupted circadian rhythm."
	close(strings)
}

func main() {
	strings := make(chan string)
	go generator(strings)
	go func() {
		for s := range strings {
			fmt.Printf("2-----%s\n", s)
		}
	}()

	for s := range strings {
		fmt.Printf("1----%s\n", s)
	}

	fmt.Printf("\n")
}
