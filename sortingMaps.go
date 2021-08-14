package main

import (
	"fmt"
	"sort"
)

var(
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo":
		43, "lima": 98,
	}
)




func main()  {
	fmt.Printf("%+v",barVal)
	fmt.Println("unsorted:")

	//for k, v:= range barVal{
	//	fmt.Printf("key: %v, value: %v \n", k,v)
	//}


	//storing all keys in separate slice
	keys := make([]string, len(barVal))
	i := 0
	for k,_ := range barVal{
		//barVal[k] = v
		keys[i] =k
		i++
	}


	sort.Strings(keys)
	newMap := make(map[string]int)
	for _, key := range keys {
		//fmt.Printf("key: %v, value: %v \n ", key, barVal[key])
		newMap[key] = barVal[key]
	}



	fmt.Printf("%+v",newMap)

	//sort.Strings(keys)
	//fmt.Println()
	//fmt.Println("\n sorted:d")
	//
	//for _,k := range keys {
	//	fmt.Printf("key: %v, value: %v /",k,barVal[k])
	//}
}