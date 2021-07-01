package pkg

import "fmt"

type SpecificPrototype struct {
	Id   int
	Name string
}

func NewSpecificPrototype(id int, name string) *SpecificPrototype {
	return &SpecificPrototype{id, name}
}

func (s *SpecificPrototype) specificAction() {
	fmt.Printf("SpecificPrototype %v is specificAction\n", *s)
}
