package main

import (
	"fmt"

	"stathat.com/c/consistent"
)

func main() {
	clientsCircle := consistent.New()

	clientsCircle.Add("111111")
	clientsCircle.Add("222222")
	clientsCircle.Add("333333")

	fmt.Println(clientsCircle.Members())

	clientsCircle.Set([]string{"44", "55", "66"})

	fmt.Println(clientsCircle.Members())

	s, _ := clientsCircle.Get("leslie")
	fmt.Println("1----", s)

	s, _ = clientsCircle.Get("kurt")
	fmt.Println("2----", s)

	clientsCircle.Remove("toy")

	s, _ = clientsCircle.Get("poul")
	fmt.Println("3----", s)

	s2, s3, _ := clientsCircle.GetTwo("leslie")
	fmt.Println("4----", s2, s3)
}
