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

func RangeAndChange() {
	nums := []int{1, 2, 3, 4, 5, 6} // [6]int{1, 2, 3, 4, 5, 6}
	li := len(nums) - 1
	for i, v := range nums { // 数组 vs 切片
		if i == li {
			nums[0] += v
		} else {
			nums[i+1] += v
		}
	}
	fmt.Println(nums)
}

// ForBreak break只会跳出当前逻辑块
func ForBreak() {
	for j := 0; j < 10; j++ {
		i := 0
		for {
			i++
			fmt.Println(i)
			if i == 10 {
				break
			}
		}
		fmt.Printf("task %d finished, i = %d", j, i)
	}
	fmt.Printf("task finished")
}

// ForBreakTag 跳到Loop后，不会再进入Loop下的逻辑块
func ForBreakTag() {
	i := 0
Loop:
	for {
		i++
		fmt.Println(i)
		if i == 10 {
			break Loop
		}
	}
	// fmt.Printf("task finished, i = %d", i)
}
