package main

// Human ...
type Human interface {
	sayHi()
}

// NewGroupFactory ...
func NewGroupFactory(name string) *Group {
	return &Group{
		Name: name,
	}
}

// Group ...
type Group struct {
	Name    string
	members []Human
}

// AddMember ...
func (g *Group) AddMember(human Human) {
	g.members = append(g.members, human)
}

// ListMembers ...
func (g *Group) ListMembers() {
	for _, member := range g.members {
		member.sayHi()
	}
}
