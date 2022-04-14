package error

import "fmt"

func DivideWithoutRecover(a, b int) int {
	return a / b
}

func Divide(a, b int) (ret int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic happened! a: %d, b: %d\n", a, b)
			// defer好像没法通过return返回值，如果没有赋值给ret，走到这里返回的就是零值
			//return -1
			ret = -1
		}
	}()
	return a / b
}
