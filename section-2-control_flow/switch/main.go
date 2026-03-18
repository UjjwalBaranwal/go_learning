package main

import (
	"fmt"
	"time"
)

func main(){
	var day string = "Monday"
	fmt.Println("Today is ",day)
	switch day{
	case "Sunday","Saturday":
		fmt.Println("Weekend !! YaaaaY")
	case "Mondat","Tuest":
		fmt.Println("Too Much Work")
	default:
		fmt.Println("I am sleepy")
	}

	switch hour:=time.Now().Hour();{
	case hour < 12:
		fmt.Println("Good Morning")
	case hour <17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	checkType := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("this is INT")
		case string:
			fmt.Println("This is String")
		case bool:
			fmt.Println("This is BOOL")
		}
	}

	checkType(12)
	checkType("HEY_UJJWAL")
}