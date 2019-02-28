# The Start

```golang

package main

func main() {

	group := NewGroupFactory("Gang")

	candidates := []string{"Spectral", "GopherOne", "GopherJoe"}
	newMember := Wannabee{}

	// let fly }>Y<{
	for _, name := range candidates {
		newMember.Name = name
		group.AddMember(newMember)
	}

	group.ListMembers()

	MajorFunctionUpgrade()
}


```

