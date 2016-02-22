package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "worker_121212"

	fmt.Println("HasPrefix---", strings.HasPrefix(s, "worker_"))
	fmt.Println("TrimLeft---", strings.TrimLeft(s, "ow"))
	fmt.Println("TrimPrefix---", strings.TrimPrefix(s, "ow"))
}
