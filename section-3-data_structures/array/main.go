package main

import "fmt"

func main() {
	var number [2]int
	number[0] = 1
	number[1] = 2
	fmt.Println(number[0])
	// number[2] // out of bound error
	
	fmt.Printf("%+v",number)
	fmt.Println()
	primes:=[5]int{2,3,5,7,11}
	fmt.Printf("%+v",primes)
	fmt.Println()
	
	for i:=0;i<len(primes);i++{
		fmt.Printf("%d ,",primes[i])
	}
	fmt.Println()

	// 2d matrix
	var mat [2][3]int // default value 0
	mat[0][0]=1
	fmt.Printf("%+v",mat)

}