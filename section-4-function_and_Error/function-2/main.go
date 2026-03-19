package main

import "fmt"

func fact(i int) int {
	if i <= 1 {
		return 1
	}
	return fact(i-1) * i
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(fact(5))
	nextSeq := intSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("hey")
}
