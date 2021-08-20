package main
import(
	"fmt"
)

type Boat struct {
	Name string
	occupants []string
	done bool

}

func (b *Boat) AddOccupant(name string) *Boat {
	b.occupants = append(b.occupants,name)
	return b
}

func (b *Boat) Remove(id int) *Boat {
	//b.occupants = b.occupants[id:]
	id -= 1
	b.occupants = append(b.occupants[:id],b.occupants[id+1:]...)
	return b
}

func (b *Boat) isDone(id int) *Boat {
	if b.done {
		id -= 1
		b.occupants = append(b.occupants[:id],b.occupants[id+1:]...)

	}else{
		fmt.Println("The",b.Name, "has the following occupants :")
		for i, num := range b.occupants{
			fmt.Printf("%v: %s \t",i+1, num)
		}
	}
	return b
}

func (b Boat) Manifest()  {

	fmt.Println("The",b.Name, "has the following occupants :")
	for i, num := range b.occupants{
		fmt.Printf("%v: %s \t",i+1, num)
	}
}




var b = &Boat{
	Name:      "S.S Tambarie",
	occupants: nil,
	done: false,
}

func main()  {
	b.AddOccupant("Sammy the shark")
	b.AddOccupant("Tammy gets")
	b.AddOccupant("clean the house")
	b.AddOccupant("arrange my bathroom")
	b.Manifest()
	//b.Remove(4)
	b.isDone(2)
	b.Manifest()

}