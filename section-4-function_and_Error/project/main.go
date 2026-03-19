package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Op      string
	InputA  int
	InputB  int
	Message string
}

const (
	division         = "Division"
	divisionErrorMsg = "division by zero"
)

func (e MathError) Error() string {
	inputs := []string{}
	if e.Op == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}
	return fmt.Sprintf("Math Error in %s (%s): %s",
		e.Op,
		strings.Join(inputs, ","),
		e.Message,
	)

}

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Op:      division,
			InputA:  a,
			InputB:  b,
			Message: divisionErrorMsg,
		}
	}
	return a / b, nil
}

func doSum(nums ...int) int {
	sum := 0
	for _, it := range nums {
		sum += it
	}
	return sum
}

func main() {
	fmt.Println(doSum(1, 2, 3, 4))
	fmt.Println(doSum(1, 2))
	value, err := safeDivision(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("SafeDivision", value)
}
