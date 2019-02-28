package main

import "fmt"
import (
	hellomodV2 "github.com/donvito/hellomod/v2"
)

func main() {

	hellomodV2.SayHello("Melvin")

	group := NewGroupFactory("Gang")
	group.AddMember(&Wannabee{"Kurt"})
	group.ListMembers()

	MajorFunctionUpgrade()
}

// Wannabee ...
type Wannabee struct {
	Name string
}

func (w *Wannabee) sayHi() {
	fmt.Println(`"Hey!", says the wannabee called`, w.Name)
}

// MajorFunctionUpgrade ...
func MajorFunctionUpgrade() {
	fmt.Println("Modules are SWELL!!!")
}
