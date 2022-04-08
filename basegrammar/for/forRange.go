package _for

import "fmt"

func PrintNumberByRange(numbers []string) {
	for i, num := range numbers {
		fmt.Println(i, num)
	}
}

func PrintNumberByIndex(numbers []string) {
	length := len(numbers)
	for i := 0; i < length; i++ {
		fmt.Println(i, numbers[i])
	}
}
