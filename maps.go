package main

import "fmt"

func main()  {
	//var x map[string]int
	/*
	x := make(map[string]int)

	x["key"] = 23
	fmt.Println(x)
	delete(x,"key")

	fmt.Println(x)
	 */

	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"
	//elements["Li"] = "Lithium"
	//elements["Be"] = "Beryllium"
	//elements["B"] = "Boron"
	//elements["C"] = "Carbon"
	//elements["N"] = "Nitrogen"
	//elements["O"] = "Oxygen"
	//elements["F"] = "Fluorine"
	//elements["Ne"] = "Neon"
	//
	////fmt.Println(elements["Mg"])
	//name, ok := elements["C"]
	//fmt.Println(name,ok)

	//elements := map[string]map[string]string{
	//	"H":{
	//		"name":"Hydrogen",
	//		"state":"gas",
	//	},
	//	"He":{
	//		"name":"Helium",
	//		"state":"gas",
	//	},
	//}
	//
	//fmt.Println(elements)
	//
	//if el,ok := elements["Li"]; ok{
	//	fmt.Println(el["name"],el["state"])
	//}else{
	//	println("not in list")
	//}

	//Using he MAKE keyword
	/*
	m := make(map[string]int)
	m["tambarie"] = 8037700350
	*/
	// Using the map literal
	/*
	k := map[string]int{
		"emmanuel":90,
		"tambarie":100,
	}

	for _,value := range k{
		fmt.Printf("%d \n",value)
	}
	*/


	//fmt.Printf("%v\n",m)
	//fmt.Printf("%v\n",k)
	//
	//if value,isPresent := m["tambarie"]; isPresent{
	//	fmt.Printf("%v %v\n",value,isPresent)
	//}
	//
	//delete(m,"tambarie")
	//fmt.Printf("%v\n",m)


	items := make([]map[int]int,5)
	for i := range items {
		items[i] = make(map[int]int,1)
		items[i][1] = 2
	}
	fmt.Println(items)
}