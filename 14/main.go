package main

import "fmt"

var sls1 = []string{
	"cat",
	"cat",
	"dog",
	"tree",
	"dog",
	"世界",
	"рыба",
	"⌘",
	"⌘",
	"рыба",
}

type mySet map[string]struct{}

func NewSet() mySet {
	return make(map[string]struct{})
}

func (s mySet) Insert(str string) bool {
	if _, ok := s[str]; ok {
		return false
	} else {
		s[str] = struct{}{}
	}
	return true
}

func (s mySet) String() string {
	var keys string = "["
	var size int
	for key, _ := range s {
		keys += key
		size++
		if size != len(s) {
			keys += " "
		}
	}
	keys += "]"
	return keys
}

func main() {
	set := NewSet()
	for _, s := range sls1 {
		f := set.Insert(s)
		if f == false {
			fmt.Printf("%v is already contains\n", s)
		}
	}
	fmt.Println("Set: ", set)
}
