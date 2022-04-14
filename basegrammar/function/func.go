package function

func defineFunc() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

// 只要参数相同、返回值相同，则函数都可以作为参数传递进来
func useFunc(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func Calculate(a, b, c, d int) int {
	add := defineFunc()
	minus := func(a, b int) int {
		return a - b
	}
	res := useFunc(a, b, add)
	res = minus(res, c) // res = useFunc(a, b, minus)
	// 匿名函数
	// 如果不是通过参数传递，而是函数内部直接访问外部变量，则称为闭包
	return func(res, d int) int {
		return res * d
	}(res, d)
}
