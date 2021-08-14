package main

import "fmt"

func main()  {
	x :=  []float64{
		 98,
		 93,
		 77,
		 82,
		 83,
	}
	y := make([]float64,3)

	copy(y,x)

	fmt.Println(x)
	fmt.Println(y)



	slice1 := []int{1,2,3}
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	//
	//var total float64
	//
	//for i := 0; i < len(x); i++{
	//	total += x[i]
	//}
	//
	//mean := total/float64(len(x))
	//fmt.Println(mean)

}