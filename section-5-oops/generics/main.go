package main

import "fmt"

type Number interface {
	int | float32 | float64 | string
}

func doSum[T Number](nums ...T) T {
	var total T
	for _, it := range nums {
		total += it
	}
	return total
}
func main() {
	// v := doSum[string]("a", "b", "c")
	v := doSum("a", "b", "c")
	fmt.Printf("%v\n", v)
	fmt.Printf(v)
}
