package main

import "fmt"

func main()  {
	/*
	b := new(bool)
	fmt.Printf("%v\n",*b)

	i := new(int)
	fmt.Printf("%v\n",*i)
	*/


	// INITIALIZATION

	type Rect struct {
		x,y float64
		width, height float64
	}

	rect := &Rect{0,0,100,200}
	fmt.Println(*rect)

	rect4 := &Rect{height: 100,width: 200}
	fmt.Println(*rect4)
}