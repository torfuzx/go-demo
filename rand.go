package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var guid [16]byte
	_, err := io.ReadFull(rand.Reader, guid[:16])
	if err != nil {
		return
	}
	guidStr := fmt.Sprintf("%x", guid)
	fmt.Println(guidStr)
}
