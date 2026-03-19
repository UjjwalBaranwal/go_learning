package main

import (
	"fmt"
	"os"
)

// “defer = schedule this work to run at the very end”
// “defer = schedule this work to run at the very end”

func simpleDefer() {
	// defer follow LIFO
	println("This is the starting line")
	defer println("defer 1")
	defer println("defer 2")
	println("This is the last line line")
}

func returnNum() (result int) {
	defer func() {
		result++
	}()
	return 5
}
func main() {
	simpleDefer()
	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println("Error in creaetign file")
	}
	defer file.Close()
	// mostly defer is used iin release resources
	println(returnNum()) // it printing 6 as defer function can interact with the named return
}
