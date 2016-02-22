package main

import (
	"fmt"
	"poit/p"

	"unsafe"
)

func main() {

	var v *p.V = new(p.V)

	var i *int32 = (*int32)(unsafe.Pointer(v))

	*i = 98

	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + 2*uintptr(unsafe.Sizeof(int32(0)))))

	*j = 763

	// v.PutI()

	// v.PutJ()

	var b []byte = []byte{'a', 'b', 'c'}

	var c *byte = &b[0]

	fmt.Println(*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(1))))

}
