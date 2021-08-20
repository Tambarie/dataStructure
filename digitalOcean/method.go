package main
import "fmt"

type Creature struct {
	name string
	greeting string
}

func (c Creature) Greet()  {
	fmt.Printf("%v, my name is %v",c.greeting, c.name)
}



var tottle = Creature{
	name: "touch",
	greeting: "hello everyone",
}

func main()  {
	tottle.Greet()
}