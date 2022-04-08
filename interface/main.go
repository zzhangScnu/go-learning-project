package main

import (
	"fmt"
	"github.com/go-learning-project/interface/cat"
)

/*
main包下推荐只放main.go入口文件（除了test文件）。当有其他文件，且在main.go文件中引用时，执行go run main.go会提示undefined。
只有非main包的依赖才会自动加载。
所以go run的时候需要将main包下的所有文件都作为参数加在命令行中。
*/
func main() {
	oneCat := cat.NewCat("yellow", 10.0, 10.0)
	showDetail(oneCat)
}

func showDetail(animal Animal) {
	fmt.Println(animal.Color(), animal.Weight(), animal.Height())
	// 类型断言（无论何时都要带上包名）
	oneCat, ok := animal.(cat.Cat)
	fmt.Println(ok)
	if ok {
		oneCat.SetName("mi")
		fmt.Println(oneCat.Name())
	}
}

type Animal interface {
	Color() string
	Height() float64
	Weight() float64
}
