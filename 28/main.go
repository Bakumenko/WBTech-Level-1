package main

import (
	"fmt"
	"task28/pkg"
)

func main() {
	client := pkg.NewClient(1, "Oleg")
	proto := pkg.NewPrototype(1, "Proto")
	spesificProto := pkg.NewSpecificPrototype(1, "Specific proto")
	adapter := pkg.NewAdapter(spesificProto)

	fmt.Println("Start proto")
	client.Doing(proto)

	fmt.Println("\nStart spec proto")
	client.Doing(adapter)
}
