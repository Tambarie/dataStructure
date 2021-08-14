package main

import (
	"fmt"
	"strconv"
)

////////////////
const LIMIT = 4
type Stack struct {
	ix	int
	data[LIMIT]int
}

func (st *Stack) Push(n int)  {
	if st.ix == LIMIT{
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) pop()  int{
	if st.ix >0{
		st.ix--
		element := st.data[st.ix]
		st.data[st.ix] = 0
		return element
	}
	return -1
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix;ix++{
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "]"
	}
	return str
}


func main()  {
	st1 := new(Stack)
	st1.Push(7)
	fmt.Printf("%v\n",st1)
	p := st1.pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n",st1)
}