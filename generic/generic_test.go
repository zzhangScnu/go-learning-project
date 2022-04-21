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

func TestIndexOf(t *testing.T) {
	fmt.Println(IndexOf([]int{1, 2, 3}, 2))
	fmt.Println(IndexOf([]float64{1.1, 2.2, 3.3}, 4.4))
}

func TestSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	nums1, nums2 := Slice(nums)
	fmt.Println(nums1)
	fmt.Println(nums2)
}
