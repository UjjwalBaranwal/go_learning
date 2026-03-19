package main

import (
	"fmt"
	"slices"
)

//slices are dynamic

func main(){
	names:=[]string{"Ujjwal","Baranwal","PRO"}// this is a slice as we dont give its length
	fmt.Println(names)

	items := make([]int,3,5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	// items[3]
	// fmt.Println(items[3])// it will give out of bound error (run time error)
	items=append(items, 2)
	fmt.Println(items[3])//now it will not give anhy error as the length increase here
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	fmt.Printf("%+v\n", items[3:7])
	//----------Advanced Slice operation
	fmt.Println("----Advanced Slice operation")

	ori:=[]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Printf("ori contains %+v values with len : %d and Capacity %d\n",ori,len(ori),cap(ori))
	s1:=ori[1:4]//[inc,exc)
	fmt.Printf("s1 contains %+v values with len : %d and Capacity %d\n",s1,len(s1),cap(s1))
	s2:=ori[:6]
	fmt.Printf("s2 contains %+v values with len : %d and Capacity %d\n",s2,len(s2),cap(s2))
	s3:=ori[3:]
	fmt.Printf("s3 contains %+v values with len : %d and Capacity %d\n",s3,len(s3),cap(s3))
	s4:=ori[:]
	fmt.Printf("s4 contains %+v values with len : %d and Capacity %d\n",s4,len(s4),cap(s4))
	a:=slices.Contains(s4,3)
	fmt.Println(a)
	s4=append(s4, 1000)
	fmt.Printf("s4 contains %+v values with len : %d and Capacity %d\n",s4,len(s4),cap(s4))
}
