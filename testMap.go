package main

import "fmt"

func main()  {



//fmt.Println(findDay(8))


	items := make([]map[int]int,5)
	for i, _ := range items {
		items[i] = make(map[int]int,1)
		items[i][1] =2

	}

	fmt.Println(items)
}
//func findDay(n int) string {
//
//	var Days = map[int]string{
//		1: "Monday",
//		2: "Tuesday",
//		3: "Wednesday",
//		4: "Thursday",
//		5: "Friday",
//		6: "Saturday",
//		7: "Sunday",
//	}
//
//	if key, isPresent :=  Days[n]; isPresent{
//		return  fmt.Sprintf("%v\n",key)
//	}else{
//		return fmt.Sprintln("Not present")
//	}
//}