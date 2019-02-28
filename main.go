package main

import "fmt"

func main() {

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
