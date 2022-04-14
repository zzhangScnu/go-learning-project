package error

import (
	"errors"
	"fmt"
)

func DivideByCheck(a, b int) (int, error) {
	if b == 0 {
		// fmt.Errorf...
		return -1, errors.New("cannot divide zero")
	}
	return a / b, nil
}

func PrintDivideResult(a, b int) {
	if r, err := DivideByCheck(a, b); err == nil {
		fmt.Println("result is: ", r)
	} else {
		fmt.Println(err.Error())
	}
}
