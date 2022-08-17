package panic

import "testing"

func TestRecoverByAnotherGoRoutine(t *testing.T) {
	recoverByAnotherGoRoutine()
}

func TestRecoverByAnotherFunction(t *testing.T) {
	recoverByAnotherFunction()
}

func TestRecoverByUnnamedFunc(t *testing.T) {
	recoverByUnnamedFunc()
}

func TestRecoverCorrectly(t *testing.T) {
	recoverCorrectly()
}
