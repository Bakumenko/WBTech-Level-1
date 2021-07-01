package pkg

import "fmt"

type Adapter struct {
	Proto *SpecificPrototype
}

func NewAdapter(adaptee *SpecificPrototype) *Adapter {
	return &Adapter{adaptee}
}

func (a *Adapter) action() {
	fmt.Println("Adapter is action")
	a.Proto.specificAction()
}
