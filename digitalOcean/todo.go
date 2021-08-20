package main
import(
	"fmt"
)

type ToDo struct {
	name string
	title string
	list []string
	done bool
}

func (t *ToDo) AddItem(todo string) *ToDo {
	t.list = append(t.list,todo)
	return t
}

func (t ToDo) ShowList()  {
	fmt.Println("Hello",t.name, "you have the following items on your ToDo List")

	for i, item := range t.list{
		fmt.Printf("%v: %v\n",i+1, item)
	}
}

func (t *ToDo) RemoveItem(id int)  *ToDo{
	id -= 1
	t.list = append(t.list[:id], t.list[1+id:]...)
	return t
}

func (t *ToDo) IsDone(id int)  *ToDo{
	if t.done{
		id -= 1
		t.list = append(t.list[:id], t.list[1+id:]...)
	}else{
		fmt.Println("Hello",t.name, "you have the following items on your ToDo List")

		for i, item := range t.list{
			fmt.Printf("%v: %v\n",i+1, item)
		}
	}
	return t
}





var td = ToDo{
	name: "emma",
	title: "To Do List",
	list: nil,
	done: true,
}

func main()  {
	td.AddItem("Clean the house")
	td.AddItem("Arrange the house")
	td.AddItem("Wash the house")
	td.AddItem("Read the house")
	td.ShowList()
	td.RemoveItem(1)
	td.ShowList()
	td.IsDone(2)
	td.ShowList()
}