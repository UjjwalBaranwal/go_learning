package main

import (
	"errors"
	"fmt"
	"strings"
)

func greetName(name string) {
	fmt.Println("Hello ", name)
}

func add(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func calcArea(width float32, height float32) float64 {
	if width < 0 || height < 0 {
		fmt.Println("Error : invalid input")
		return 0.0
	}
	return float64(width) * float64(height)
}

func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero error")
	}
	return float64(a) / float64(b), nil
}

func splitName(fullName string) (string, string) {
	parts := strings.Split(fullName, " ")
	firstName := parts[0]
	lastName := parts[1]
	return firstName, lastName
}
func main() {
	greetName("Ujjwal")
	add(1, 2)
	area := calcArea(1.324, 1.456)
	fmt.Printf("Calc area : %.2f\n", area)

	divideAndLog := func(a int, b int) {
		value, err := divide(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
	divideAndLog(4, 2)
	fn, ln := splitName("Ujjwal BAranwal")
	fmt.Println(fn, ln)
}
