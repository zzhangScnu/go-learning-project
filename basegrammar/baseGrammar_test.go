package basegrammar

import (
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
