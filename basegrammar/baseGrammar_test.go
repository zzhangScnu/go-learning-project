package basegrammar

import (
	"fmt"
	_for "github.com/go-learning-project/basegrammar/for"
	"github.com/go-learning-project/basegrammar/function"
	_if "github.com/go-learning-project/basegrammar/if"
	_switch "github.com/go-learning-project/basegrammar/switch"
	"testing"
)

func TestSwitch(t *testing.T) {
	_switch.Judge(100)
	_switch.Judge("100")
	_switch.Judge(100.0)
}

func TestFallthrough(t *testing.T) {
	_switch.Fallthrough(0)
}

func TestIf(t *testing.T) {
	_if.Judge(100)
	_if.Judge("100")
	_if.Judge(100.0)
}

func TestRangeStruct(t *testing.T) {
	_for.RangeStruct()
}

func TestRangeStructConcurrently(t *testing.T) {
	_for.RangeStructConcurrently()
}

func TestRangeString(t *testing.T) {
	_for.RangeString()
}

func TestRangeMap(t *testing.T) {
	for i := 0; i < 5; i++ {
		_for.RangeMap()
		fmt.Println("=====================")
	}
}

func TestFunc(t *testing.T) {
	r := function.Calculate(1, 2, 3, 4)
	if r != 0 {
		t.Errorf("wrong result: %d", r)
	}
}

func TestInitialize(t *testing.T) {
	function.Show()
}
