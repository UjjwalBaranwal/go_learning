package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	fmt.Printf("%+v\n",[]int{1,2,3})
	var t any

	fmt.Printf("%+v\n",t)
	// intialising the variable
	var greeting string // zero-value is an empty string
	greeting="Hello world"
	fmt.Println(greeting)
	var count int
	count=10
	fmt.Println(count)
	var firstName,lastName string
	firstName="Ujjwal"
	lastName="Baranwal"
	fmt.Println(firstName,lastName)
	fmt.Printf("My name is %s %s\n",firstName,lastName)
	var isOwn bool
	isOwn=false
	fmt.Println(isOwn)

	//variable declare and initilaise
	email:="Ub" // infer type

	fmt.Println(email)


 	 	}
