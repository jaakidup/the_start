package main

import "fmt"

// Wannabee ...
type Wannabee struct {
	Name string
}

func (w Wannabee) sayHi() {
	fmt.Println(`"Hey!", says the wannabee called`, w.Name)
}

// MajorFunctionUpgrade ...
func MajorFunctionUpgrade() {
	fmt.Println("Modules are SWELL!!!")
}
