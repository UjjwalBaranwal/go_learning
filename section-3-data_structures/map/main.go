package main

import "fmt"

func main() {
	studentGrade := map[string]int{
		"Ujjwal":   90,
		"Baranwal": 100,
		"PRO":      101,
	}
	fmt.Printf("%+v\n",studentGrade)
	studentGrade["Alice"]=10
	fmt.Printf("%+v\n",studentGrade)
	alice,ok:=studentGrade["Alice"]
	if ok{
		fmt.Printf("Alice got %d marks\n",alice)
		// fmt.Printf("Alice got %+v marks\n",alice)
	}
	key:="Bob"
	if value,ok:=studentGrade[key];ok{
		fmt.Printf("%s got %d marks\n",key,value)
	}else{
			fmt.Printf("%s not exist in map\n",key)
	}
	
	delete(studentGrade,"Alice")//it will delete from the map
	fmt.Printf("%+v\n",studentGrade)

	configs:=make(map[string]string)

	fmt.Printf("%+v %T\n",configs,configs)
	if configs == nil{
		fmt.Println("Config is empty")
	}
}