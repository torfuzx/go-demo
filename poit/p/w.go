package p

import (
	"fmt"

	"unsafe"
)

type W struct {
	b byte

	i int32

	j int64
}

func init() {

	var w *W = new(W)

	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))

}
