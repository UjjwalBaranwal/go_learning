// A function is variadic if its last parameter uses ...
// rules of Varidadic function
/*
1. Only one variadic parameter allowed
2. Must be last parameter
*/

package main

import "fmt"

func doSum(nums ...int){//variadic function
	sum:=0
	for _,it := range nums{
		sum+=it
	}
	fmt.Println("total_sum : ",sum)
}

func multMyA(a int,b ...int){
	for _,it := range b{
		fmt.Printf("%d ",it*a)
	}
	println()
}

func main(){
	doSum(1,2,3)
	doSum(1,2)
	doSum(1,2,3,4,5)

	multMyA(2,2,3,4,5,6)
	nums:=[]int{1,2,3,4}
	multMyA(2,nums...)
}