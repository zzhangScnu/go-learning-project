package basegrammar

import (
	_switch "github.com/go-learning-project/basegrammar/switch"
	"testing"
)

func TestSwitch(t *testing.T) {
	_switch.Judge(100)
	_switch.Judge("100")
	_switch.Judge(100.0)
}
