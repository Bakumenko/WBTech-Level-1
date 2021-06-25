package main

import (
	"encoding/binary"
	"fmt"
)

var (
	i      int   = 4
	number int64 = 107
)

func main() {
	if i > 0 && i < 8 {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(number))

		fmt.Printf("Old nubmer: %v\n", number)
		fmt.Printf("Old nubmer in bytes: %v\n", b)

		b[i] = 1
		newNubmer := int64(binary.LittleEndian.Uint64(b))

		fmt.Printf("New nubmer in bytes: %v\n", b)
		fmt.Printf("New number: %v\n", newNubmer)
	} else {
		fmt.Printf("%v is out of range 0..7", i)
	}
}
