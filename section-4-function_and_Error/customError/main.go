package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrNumTooLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string { // we have to uuse this whenever we are usign custom error
	return op.Message
}

func newOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	if a > 1000 {
		return 0, ErrNumTooLarge
	}
	return float64(a) / float64(b), nil
}

func main() {

	divideAndLog := func(a int, b int) {
		val, err := divide(a, b)
		if err == nil {
			fmt.Printf("%d / %d = %.2f\n", a, b, val)
		} else if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("Divide by zero error")
		} else if errors.Is(err, ErrNumTooLarge) {
			fmt.Println("Number is too big")
		}
	}

	divideAndLog(1, 2)
	divideAndLog(1, 0)
	divideAndLog(10000000, 0)
}
