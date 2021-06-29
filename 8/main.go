package main

import (
	"encoding/binary"
	"fmt"
)

var (
	index  int   = 1
	number int64 = 1
)

func main() {
	if index > 0 && index < 8 {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(number))

		fmt.Printf("Old number: %v\n", number)
		fmt.Printf("Old number in bytes: %v\n", b)

		b[index] = 1
		newNubmer := int64(binary.LittleEndian.Uint64(b))

		fmt.Printf("New number in bytes: %v\n", b)
		fmt.Printf("New number: %v\n", newNubmer)
	} else {
		fmt.Printf("%v is out of range 0..7", index)
	}
}
