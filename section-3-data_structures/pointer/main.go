package main

import "fmt"

// in go , pointers not perform arithmetic operation
func modifyVal(val int){
	val*=10
	fmt.Printf("Modified through modifyVal Function : %d\n",val)
}

func modifyPtr(val *int){
	if val==nil{
		fmt.Println("value is nil")
		return
	}
	*val=*val*10//de-refrencing the val
	fmt.Printf("Modified value through modifyPtr %d\n",*val)
}
func main(){
	age:=10
	fmt.Printf("Value of the age : %d\n",age)
	fmt.Printf("Address of the age : %d\n",&age)
	fmt.Printf("Address of the age : %x\n",&age)
	agePtr:=&age // address of the age
	fmt.Printf("Address of the age : %x\n",agePtr)
	*agePtr=11
	fmt.Printf("Value of the age after changing through ptr : %d\n",age)
	fmt.Printf("Value of the age after changing through ptr : %d\n",*agePtr)
	val:=20
	modifyVal(val)
	println(val)
	modifyPtr(&val)
	println(val)

	grade:=10
	gradePtr:=&grade // this store the location of the grade
	// this create a new variable and store the value of grade
	// gradePtr have a new addres
	fmt.Printf("address of grade %x\n",&grade)
	fmt.Printf("address of grade stored in gradePtr %x\n",gradePtr)
	fmt.Printf("address of the gradePtr : %x\n",&gradePtr)
	fmt.Printf("Value hold by the address of the grade ptr : %x",*(&gradePtr))
}