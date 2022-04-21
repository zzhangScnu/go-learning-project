package generic

import (
	"fmt"
	"testing"
)

func TestParam(t *testing.T) {
	if AsParam(1, 2) != 3 {
		t.Errorf("test failed!")
	}
}

func TestInterface(t *testing.T) {
	if AsInterface(1, 2) != 3 {
		t.Errorf("test failed!")
	}
}

func TestIntArray(t *testing.T) {
	result := AsArray([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("test failed!")
	}
	fmt.Println(result)
}

func TestFloatArray(t *testing.T) {
	result := AsArray([]float64{1.1, 2.2, 3.3})
	if result != 6.6 {
		t.Errorf("test failed!")
	}
	fmt.Println(result)
}
