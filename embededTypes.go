package main

import (
	"fmt"
	"math"
)

// Point //////////////////
type Point struct {
	x,y float64
}

// NamedPoint ///////////////////
type NamedPoint struct {
	Point
	name string
}

// Abs ///////////////////////////
func (a *Point) Abs() float64 {
	calc := math.Sqrt(a.x*a.x + a.y*a.y)
	return calc
}

// Log //////////////////////////////
///////////////////////////////
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}



func main() {
	e := new(Customer)
	e.Name = "Barack Obama"
	e.log = new(Log)
	e.log.msg = "1 - Yes we can !"
	//shorter
	f := &Customer{"Michelle Obama",&Log{"2-Yes we can"}}
	fmt.Println(f.log)
	f.Log().Add("2 -After me, the world will be a better place")





	//////////////////////////////////////////////////
	np := &NamedPoint{Point{3,4},"tambarie"}
	d := np.Abs()

	absolute :=  &Point{4,5}
	c := absolute.Abs()
	fmt.Println(c)
	fmt.Println(d)
	/////////////////////////////////////////////////
}

func (l *Log) Add(s string)  {
	l.msg +="\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log  {
	return c.log
}