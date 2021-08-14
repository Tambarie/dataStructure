package main

import (
	"fmt"
	"math"
)

// Declaring  a struct
type Circle struct {
	radius float64
}

// func areaOfCircle(c Circle) float64 {
// 	area := math.Pi * c.radius * c.radius
// 	return   math.Floor(area) 
// }


func (c Circle) area() float64 {
	area := math.Pi * c.radius*c.radius
	return area
	
}

func main() {
	var circle Circle

	circle = Circle{radius: 20}
	fmt.Println(circle.area())
}
