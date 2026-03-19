package main

import "fmt"

// error = “handle this”
// panic = “this should never happen”

// when panic happen , the whole program crashed and only the defer funciton are executed

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("I am panicing")
	}
	fmt.Println("Exceuted without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic ", r)
		}
	}()
	mightPanic(true)
}
func main() {
	// panic("Oh shit")
	recoverable()
}
