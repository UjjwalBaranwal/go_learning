package main

import "fmt"

func main(){
	// for -- only way to loop
	fmt.Println("----------------------------")
	fmt.Println("-------For Loop Start--------------")
	for i:=1;i<5;i++{
		fmt.Println(i)
	}
	fmt.Println("-------For Loop End--------------")
	fmt.Println("----------------------------")
	fmt.Println("-------While Loop Start--------------")
	//just mimicking the while nature 
	// no keywords like while , do-while is in go
	k:=10
	for k>=0{
		fmt.Println(k)
		k--
	}
	fmt.Println("-------While Loop End--------------")
	// similarily we can do infinite loop
	// ini:=1
	// for ini>=1{
	// 	fmt.Println("Go is Gooooo")
	// }
	fmt.Println("--------------------------")
	fmt.Println("-------Print Even Number----------------")
	for i:=0;i<=10;i++{
		if i&1 == 1{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("--------------------------")
	
	fmt.Println("-----Array Iterating-----------")
	items := [3]string{"Go","Java","Python"}
	for index,value := range items{
		fmt.Printf("index : %d , value : %s\n",index,value)
	}
}