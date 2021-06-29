package main

import "fmt"

var sls2 = []string{
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

type mySet2 struct {
	keys []string
}

func NewSet2() mySet2 {
	return mySet2{[]string{}}
}

func (s *mySet2) Insert(str string) bool {
	for _, val := range s.keys {
		if val == str {
			return false
		}
	}
	s.keys = append(s.keys, str)
	return true
}

func (s mySet2) String() string {
	var keys string = "["
	var size int
	for _, key := range s.keys {
		keys += key
		size++
		if size != len(s.keys) {
			keys += " "
		}
	}
	keys += "]"
	return keys
}

func main() {
	set := NewSet2()
	for _, s := range sls2 {
		f := set.Insert(s)
		if f == false {
			fmt.Printf("%v is already contains\n", s)
		}
	}
	fmt.Println("Set: ", set)
}
