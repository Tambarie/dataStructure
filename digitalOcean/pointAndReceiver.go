package main
import (
	"fmt"
)

type Submersible  interface {
	Dive()
}

type Shark struct {
	name string
	isUnderwater bool
}

func (s Shark) String() string {
	if s.isUnderwater{
		return fmt.Sprintf("%s is underwater", s.name)
	}
	return fmt.Sprintf("%s is on the surface", s.name)
}

func (s *Shark) Dive()  {
	s.isUnderwater = true
}

func submerge(s Submersible)  {
	s.Dive()
}

func main()  {
	s := &Shark{
		name: "tambari",
	}
	fmt.Println(s)
	submerge(s)
	fmt.Println(s)
}

