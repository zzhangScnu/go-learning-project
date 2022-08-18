package basegrammar

import (
	"fmt"
	_for "github.com/go-learning-project/basegrammar/for"
	"github.com/go-learning-project/basegrammar/function"
	_if "github.com/go-learning-project/basegrammar/if"
	string2 "github.com/go-learning-project/basegrammar/string"
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

func TestRangeAndChange(t *testing.T) {
	_for.RangeAndChange()
}

func TestNums(t *testing.T) {
	fmt.Println("run ConvertToIntByForRange...")
	string2.ConvertToIntByForRange()
	fmt.Println("run ConvertToIntByForI...")
	string2.ConvertToIntByForI()
	fmt.Println("run ConvertToIntByRune...")
	string2.ConvertToIntByRune()
}

func TestString(t *testing.T) {
	fmt.Println("run RangeStringByI...")
	string2.RangeStringByI()
	fmt.Println("run RangeString...")
	string2.RangeString()
	fmt.Println("run RangeStringRune...")
	string2.RangeStringRune()
	fmt.Println("run RangeByte...")
	string2.RangeByte()
	fmt.Println("run Compare...")
	string2.Compare()
}

func TestBreak(t *testing.T) {
	_for.ForBreak()
}

func TestBreakLoop(t *testing.T) {
	_for.ForBreakTag()
}
