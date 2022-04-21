package generic

import "errors"

func AsParam[T int | float64](i T, j T) T {
	return i + j
}

type numType interface {
	int | int32 | int64 | float32 | float64
}

func AsInterface[T numType](i, j T) T {
	return i + j
}

func AsArray[T numType](nums []T) T {
	var result T
	for _, val := range nums {
		result = result + val
	}
	return result
}

func IndexOf[T comparable](candidates []T, target T) (error, int) {
	for i, val := range candidates {
		if val == target {
			return nil, i
		}
	}
	return errors.New("cannot find target"), -1
}

func Slice[T any](nums []T) ([]T, []T) {
	mid := len(nums) / 2
	return nums[:mid], nums[mid:]
}
