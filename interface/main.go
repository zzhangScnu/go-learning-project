package main

import (
	"fmt"
	"github.com/go-learning-project/interface/cat"
)

func main() {
	oneCat := cat.NewCat("yellow", 10.0, 10.0)
	showDetail(oneCat)
}

func showDetail(animal Animal) {
	fmt.Println(animal.Color(), animal.Weight(), animal.Height())
	//oneCat, err = animal.(Cat)
	//oneCat.set
}

type Animal interface {
	Color() string
	Height() float64
	Weight() float64
}
