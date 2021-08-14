package main

import "fmt"


//Rectangle
type Rectangle struct {
	length,width int
}

// perimeter
func (r *Rectangle) perimeter() int {
	cal := (r.length + r.width) * 2
	return cal
}

//square
func (r *Rectangle) square() int {
	cal := r.length * r.width
	return cal
}





func main()  {
	rect := &Rectangle{10,20}
	fmt.Println(rect.perimeter())
	fmt.Println(rect.square())
}