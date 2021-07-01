package pkg

import "fmt"

type Prototype struct {
	Id   int
	Name string
}

func NewPrototype(id int, name string) *Prototype {
	return &Prototype{id, name}
}

func (p *Prototype) action() {
	fmt.Printf("Prototype %v is action\n", *p)
}
