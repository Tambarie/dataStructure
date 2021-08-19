package main

import "fmt"

type Shaper interface {
	Area() float32
}


type Square struct {
	side float32
}
func (sq *Square) Area() float32  {
	return sq.side * sq.side
}

type Rect struct {
	length, width float32
}

func (r Rect) Area() float32 {
	return r.length * r.width
}

func main()  {
	s := &Square{5}
	r := Rect{10,20}

	shapes := []Shaper{s,r}

	for n, _ := range shapes{
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}



}