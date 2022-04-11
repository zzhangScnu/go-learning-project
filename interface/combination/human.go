package combination

import "fmt"

type Human interface {
	Say()
}

type Student struct {
	Human
	Course string
}

func (s Student) Say() {
	fmt.Println("i am a student, my course is ", s.Course)
}

type Programmer struct {
	Human
	Language string
}

func (p Programmer) Say() {
	fmt.Println("i am a programmer, my language is ", p.Language)
}
