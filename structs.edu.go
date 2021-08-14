package main

import (
	"fmt"
	"strings"
	"time"
)


//////////////////////////////////
//struct definition
type struct1 struct {
	i1 int
	f1 float32
	str string

}

// Person Passing struct as a parameter
type Person struct {
	firstName string
	lastName string
}
//////////////////////////////////
func  upPerson(p *Person) string {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)

	return fmt.Sprintf("%v %v\n",p.firstName, p.lastName)
}
///////////////////////////////////////

//Composition
type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32

	int //anonymous field
	innerS //anonymous field
}

// TwoInts ////////////////////////////////////
type TwoInts struct {
	a int
	b int
}

func (s *TwoInts) adder(params int) int {
	sum := s.a + s.b + params
	return sum
}
////////////////////////////////////////
type myTime struct {
	time.Time
}

func (t myTime) clock() string {
	return t.String()[:20]
}
func main()  {
	///////////////////
	m := myTime{time.Now()}

	fmt.Println(m.clock())




	//////////////////////////
	add := new(TwoInts)
	add.a  = 10
	add.b = 10

	fmt.Println(add.adder(3))




	//////////////////////////
//composition
	//////////////////////////
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	//With struct literal
	outer2 := outerS{6,4,5,innerS{2,4}}
	fmt.Println("outer2 is :",outer2)
	//////////////////////////





	////////////////////////
	//Size of structs
	/*size := unsafe.Sizeof(Person{})
	//fmt.Printf("%s\n",size)
	fmt.Println(size)*/

    ////////////////////////

	////////////////////////
	//struct as a value type
	/*var per1 Person
	per1.firstName = "Emmanuel"
	per1.lastName = "Gbaragbo"
	fmt.Printf("%v\n",upPerson(&per1))
	//fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
*/
	//struct as a pointer:
	/*
	per2 := new(Person)
	per2.firstName = "Clinton"
	per2.lastName = "Adedeji"
	//upPerson(per2)
	fmt.Printf("%v\n",upPerson(per2))

	//struct as a struct literal:
	per3 := &Person{"Chinonso","Okike"}
	fmt.Printf("%v\n",upPerson(per3))
	*/

	////////////////////////

	////////////////////////
	/*
	//ms := new(struct1)
	//ms.i1 = 10
	//ms.f1 = 15.5
	//ms.str = "Chris"
	ms := struct1{20,30.23,"Tambarie"}

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The int is: %f\n", ms.f1)
	fmt.Printf("The int is: %s\n", ms.str)
	*/
	//////////////////////
}