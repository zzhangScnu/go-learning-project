package error

import (
	"fmt"
	"testing"
	"time"
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

func TestErrors(t *testing.T) {
	fmt.Println(fmtSomeError())
	fmt.Println(wrapSomeError())
}

func TestGoErrors(t *testing.T) {
	err := A()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func TestName(t *testing.T) {
	sleepTimeMs := 3000
	duration := time.Duration(sleepTimeMs) * time.Millisecond
	// time.Sleep(duration)
	fmt.Println(duration.Milliseconds())
	fmt.Println(duration.Seconds())
}

func TestMap(t *testing.T) {
	mapping := make(map[int]int, 0)
	_, ok := mapping[0]
	fmt.Println(ok)
}
