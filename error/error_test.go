package error

import (
	"fmt"
	"testing"
)

func TestCustomizedError(t *testing.T) {
	w := WrongAccount{"123"}
	fmt.Println(w.Error())
}

func TestPanic(t *testing.T) {
	fmt.Println(DivideWithoutRecover(1, 0))
}

func TestRecover(t *testing.T) {
	fmt.Println(Divide(1, 1))
	fmt.Println(Divide(1, 0))
}

func TestDivideByCheck(t *testing.T) {
	PrintDivideResult(1, 1)
	PrintDivideResult(1, 0)
}
