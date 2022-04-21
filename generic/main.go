package generic

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
