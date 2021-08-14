package main

import (
	"fmt"
	"log"
	"os"
)

/*
type Circle struct{
	x float64
	y float64
	r float64
}
*/



type User struct {
	Id 		int
	Name	string
	Location	string
}

// Player type player with one additional attribute
type Player struct {
	User
	GameId		int
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi, am %s from %s", u.Name, u.Location)
}


// Job Implementing a Job struct that can also behave as a logger
type Job struct {
	Command string
	Logger *log.Logger
}

func main()  {
	job := &Job{"demo",log.New(os.Stderr,"Job:",log.Ldate)}
	job.Logger.Print("test")


	//x := Player{}
	//x.Id = 42
	//x.Name ="Matt"
	//x.Location ="LA"
	//x.GameId =90932


	var p = Player{
		User{42,
			"Matt",
			"LA",
		},
			90345,
	}

	fmt.Printf("%+v\n",p.Greetings())
	//fmt.Printf("% +v\n ",x)





	}

















	// Declaring
	/*
		var c Circle
		c = Circle{x:23.4, y: 2.0, r:4.7}
		fmt.Println(c.r)

		d := new(Circle)
		fmt.Println(*d)
		fmt.Println(circleArea(&c))
		//Calling a method
		fmt.Println((&c).areaOfC())

		//Embedded struct
		var person Person
		person = Person{Name: "Tambarie"}
		fmt.Println(person.Talk())

		//Android
		a := new(Android)
		fmt.Println(a.Person.Talk())


	*/






//Structs
/*
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

*/

//Methods
/*
func (t *Circle) areaOfC() float64 {
	area := math.Pi * t.r * t.r
	return area

}
*/

//Embedded Structs
//1
/*
type Person struct {
	Name string
}

type Android struct{
	Person
	Model string
}


func (p *Person) Talk() string {
	hello :=fmt.Sprintln("Hi, my name is,",p.Name)
	return hello
*/



//2